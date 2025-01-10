package summon

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"

	prov "github.com/cyberark/summon/pkg/provider"
	"github.com/cyberark/summon/pkg/secretsyml"
)

// SubprocessConfig is an object that holds all the info needed to run
// a Summon instance
type SubprocessConfig struct {
	Args                 []string
	Provider             string
	Filepath             string
	YamlInline           string
	Subs                 []string
	Ignores              []string
	IgnoreAll            bool
	Environment          string
	RecurseUp            bool
	ShowProviderVersions bool
	FetchSecret          SecretFetcher
}

const ENV_FILE_MAGIC = "@SUMMONENVFILE"
const SUMMON_ENV_KEY_NAME = "SUMMON_ENV"

// SecretFetcher is function signature for fetching a secret
type SecretFetcher func(string) ([]byte, error)

// RunSubprocess encapsulates the logic of fetching secrets, executing the subprocess with the secrets injected.
func RunSubprocess(sc *SubprocessConfig) (int, error) {
	var (
		secrets secretsyml.SecretsMap
		err     error
	)

	subs := convertSubsToMap(sc.Subs)

	if sc.RecurseUp {
		currentDir, err := os.Getwd()
		if err != nil {
			return 0, err
		}
		sc.Filepath, err = findInParentTree(sc.Filepath, currentDir)
		if err != nil {
			return 0, err
		}
	}

	switch sc.YamlInline {
	case "":
		secrets, err = secretsyml.ParseFromFile(sc.Filepath, sc.Environment, subs)
	default:
		secrets, err = secretsyml.ParseFromString(sc.YamlInline, sc.Environment, subs)
	}

	if err != nil {
		return 0, err
	}

	env := make(map[string]string)
	tempFactory := NewTempFactory("")
	defer tempFactory.Cleanup()

	var results []prov.Result

	// Filter out non variables
	filteredResults, filteredSecrets := filterNonVariables(secrets, &tempFactory)
	results = append(results, filteredResults...)

	// Call provider with no arguments
	resultsCh, errorsCh, cleanup := prov.CallInteractiveMode(sc.Provider, filteredSecrets)
	defer cleanup()

	// This extracts the logic of handling results from provider interactive mode
	resultsFromProvider, err := handleResultsFromProvider(resultsCh, errorsCh, filteredSecrets, &tempFactory)
	results = append(results, resultsFromProvider...)

	if err != nil {
		results = nonInteractiveProviderFallback(secrets, sc, &tempFactory)
	}

EnvLoop:
	for _, envvar := range results {
		if envvar.Error == nil {
			env[envvar.Key] = envvar.Value
		} else {
			if sc.IgnoreAll {
				continue EnvLoop
			}

			for i := range sc.Ignores {
				if sc.Ignores[i] == fmt.Sprintf("%s=%s", envvar.Key, envvar.Value) {
					continue EnvLoop
				}
			}
			return 0, fmt.Errorf("Error fetching variable %v: %v", envvar.Key, envvar.Error.Error())
		}
	}

	// Append environment variable if one is specified
	if sc.Environment != "" {
		env[SUMMON_ENV_KEY_NAME] = sc.Environment
	}

	setupEnvFile(sc.Args, env, &tempFactory)

	var e []string
	for k, v := range env {
		e = append(e, fmt.Sprintf("%s=%s", k, v))
	}

	err = runSubcommand(sc.Args, append(os.Environ(), e...))
	if err != nil {
		return returnStatusOfError(err)
	}

	return 0, nil
}

func filterNonVariables(secrets secretsyml.SecretsMap, tempFactory *TempFactory) ([]prov.Result, secretsyml.SecretsMap) {
	filteredSecrets := make(secretsyml.SecretsMap)
	results := []prov.Result{}

	for key, spec := range secrets {
		if spec.IsVar() {
			filteredSecrets[key] = spec
		} else {
			k, v := formatForEnv(key, spec.Path, spec, tempFactory)
			result := prov.Result{k, v, nil}
			results = append(results, result)
		}
	}

	return results, filteredSecrets
}

func handleResultsFromProvider(resultsCh chan prov.Result, errorsCh chan error,
	filteredSecrets secretsyml.SecretsMap, tempFactory *TempFactory) (results []prov.Result, err error) {
	for {
		select {
		case result, ok := <-resultsCh:
			if !ok {
				return results, nil
			}

			spec := filteredSecrets[result.Key]

			// Set a default value if the provider didn't return one for the item
			if result.Value == "" && spec.DefaultValue != "" {
				result.Value = spec.DefaultValue
			}
			k, v := formatForEnv(result.Key, result.Value, spec, tempFactory)
			result = prov.Result{k, v, nil}
			results = append(results, result)

		// Fallback to the old implementation if either provider doesn't support interactive mode or an error occured
		case err = <-errorsCh:
			return nil, err
		}
	}
}

func nonInteractiveProviderFallback(secrets secretsyml.SecretsMap, sc *SubprocessConfig, tempFactory *TempFactory) []prov.Result {
	results := make(chan prov.Result, len(secrets))
	var wg sync.WaitGroup

	for key, spec := range secrets {
		wg.Add(1)
		go func(key string, spec secretsyml.SecretSpec) {
			var value string
			if spec.IsVar() {
				valueBytes, err := sc.FetchSecret(spec.Path)
				if err != nil {
					results <- prov.Result{key, "", err}
					wg.Done()
					return
				}
				value = string(valueBytes)
			} else {
				// If the spec isn't a variable, use its value as-is
				value = spec.Path
			}

			// Set a default value if the provider didn't return one for the item
			if value == "" && spec.DefaultValue != "" {
				value = spec.DefaultValue
			}

			k, v := formatForEnv(key, value, spec, tempFactory)
			results <- prov.Result{k, v, nil}
			wg.Done()
		}(key, spec)
	}
	wg.Wait()
	close(results)

	resultsSlice := make([]prov.Result, 0, len(secrets))
	for result := range results {
		resultsSlice = append(resultsSlice, result)
	}
	return resultsSlice
}

func returnStatusOfError(err error) (int, error) {
	if eerr, ok := err.(*exec.ExitError); ok {
		if ws, ok := eerr.Sys().(syscall.WaitStatus); ok {
			if ws.Exited() {
				return ws.ExitStatus(), nil
			}
		}
	}
	return 0, err
}

// formatForEnv returns a string in %k=%v format, where %k=namespace of the secret and
// %v=the secret value or path to a temporary file containing the secret
func formatForEnv(key string, value string, spec secretsyml.SecretSpec, tempFactory *TempFactory) (string, string) {
	if spec.IsFile() {
		fname := tempFactory.Push(value)
		value = fname
	}

	return key, value
}

func joinEnv(env map[string]string) string {
	var envs []string
	for k, v := range env {
		envs = append(envs, fmt.Sprintf("%s=%s", k, v))
	}

	// Sort to ensure predictable results
	sort.Strings(envs)

	return strings.Join(envs, "\n") + "\n"
}

// findInParentTree recursively searches for secretsFile starting at leafDir and in the
// directories above leafDir until it is found or the root of the file system is reached.
// If found, returns the absolute path to the file.
func findInParentTree(secretsFile string, leafDir string) (string, error) {
	if filepath.IsAbs(secretsFile) {
		return "", fmt.Errorf(
			"file specified (%s) is an absolute path: will not recurse up", secretsFile)
	}

	for {
		joinedPath := filepath.Join(leafDir, secretsFile)

		_, err := os.Stat(joinedPath)

		if err != nil {
			// If the file is not present, we just move up one level and run the next loop
			// iteration
			if os.IsNotExist(err) {
				upOne := filepath.Dir(leafDir)
				if upOne == leafDir {
					return "", fmt.Errorf(
						"unable to locate file specified (%s): reached root of file system", secretsFile)
				}

				leafDir = upOne
				continue
			}

			// If we have an unexpected error, we fail-fast
			return "", fmt.Errorf("unable to locate file specified (%s): %s", secretsFile, err)
		}

		// If there's no error, we found the file so we return it
		return joinedPath, nil
	}
}

// scans arguments for the magic string; if found,
// creates a tempfile to which all the environment mappings are dumped
// and replaces the magic string with its path.
// Returns the path if so, returns an empty string otherwise.
func setupEnvFile(args []string, env map[string]string, tempFactory *TempFactory) string {
	var envFile = ""

	for i, arg := range args {
		idx := strings.Index(arg, ENV_FILE_MAGIC)
		if idx >= 0 {
			if envFile == "" {
				envFile = tempFactory.Push(joinEnv(env))
			}
			args[i] = strings.Replace(arg, ENV_FILE_MAGIC, envFile, -1)
		}
	}

	return envFile
}

// convertSubsToMap converts the list of substitutions passed in via
// command line to a map
func convertSubsToMap(subs []string) map[string]string {
	out := make(map[string]string)
	for _, sub := range subs {
		s := strings.SplitN(sub, "=", 2)
		key, val := s[0], s[1]
		out[key] = val
	}
	return out
}
