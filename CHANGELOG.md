# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Fix: bring back the --up flag [#254 --up is broken: open secrets.yml: no such file or directory](https://github.com/cyberark/summon/issues/254).

## [0.10.1] - 2024-08-14

### Changed
- Fixed a bug that caused Summon to not properly handle a situation where a secret
  to be fetched did not exist (CNJR-6093)

## [0.10.0] - 2024-07-29

### Added
- Introduced secrets streaming from providers (CNJR-3580)
- Upgrade Go to 1.22.4
- Added NOTICES.txt file to project

## [0.9.8] - 2024-04-25

### Security
- Upgrade golang.org/x/net to v0.24.0 (CONJSE-1862)

## [0.9.7] - 2024-03-21

### Security
- Upgrade Go to 1.22 (CONJSE-1842)

## [0.9.6] - 2023-06-14

### Security
- Upgrade golang.org/x/net to v0.10.0, golang.org/x/crypto to v0.9.0,
  golang.org/x/sys to v0.8.0, and Go to 1.20
  [cyberark/summon#247](https://github.com/cyberark/summon/pull/247)
- Upgrade golang.org/x/net to v0.7.0 for CVE-2022-41721 and CVE-2022-41722 (not vulnerable)
  [cyberark/summon#245](https://github.com/cyberark/summon/pull/245)

## [0.9.5] - 2022-09-28
### Changed
- Upgraded example Dockerfile to use python:3.11
  [cyberark/summon#243](https://github.com/cyberark/summon/pull/243)
- Upgrade Go to 1.19
  [cyberark/summon#240](https://github.com/cyberark/summon/pull/240)

### Security
- Force golang.org/x/text to use v0.3.8 
  [cyberark/summon#241](https://github.com/cyberark/summon/pull/241)
- Update aruba (0.6.2 -> 2.0.0), cucumber (2.0.0 -> 7.1.0) and other necessary
  dependencies in acceptance/Gemfile.lock
  [cyberark/summon#239](https://github.com/cyberark/summon/pull/239)
- Update golang.org/x/net to v0.0.0-20220923203811-8be639271d50
  [cyberark/summon#240](https://github.com/cyberark/summon/pull/240)

## [0.9.4] - 2022-08-18
### Security
- Replaced gopkg.in/yaml.v2 v2.2.2 with v2.2.8 to address 
  SNYK-GOLANG-GOPKGINYAMLV2-1533594 and CVE-2019-11254
  [cyberark/summon#236](https://github.com/cyberark/summon/pull/236)

## [0.9.3] - 2022-06-15
### Changed
- Updated dependencies in go.mod (github.com/stretchr/testify -> 1.7.2, 
  github.com/urfave/cli -> 1.22.9, golang.org/x/net -> v0.0.0-20220607020251-c690dde0001d,
  gopkg.in/yaml.v3 -> v3.0.1)
  [cyberark/summon#234](https://github.com/cyberark/summon/pull/234)

## [0.9.2] - 2022-05-31
### Security
- Update main and acceptance base images to Golang 1.17 to fix CVE-2022-0778 and CVE-2022-1292.
  [cyberark/summon#232](https://github.com/cyberark/summon/pull/232/)

## [0.9.1] - 2021-12-22
### Changed
- Update go to 1.17 & switch to github.com/urfave/cli 
  from github.com/codegangsta/cli
  [cyberark/summon#226](https://github.com/cyberark/summon/pull/226)

## [0.9.0] - 2021-07-19

### Added
- Build for Apple M1 silicon.
  [cyberark/summon#216](https://github.com/cyberark/summon/issues/216)

### Added
- Addded portable mode for provider directory search. If now global provider directory is
  found providers are searched next to the `summon` executable in `<path_to_exe>/Providers/`
  [cyberark/summon#164](https://github.com/cyberark/summon/issues/164)

### Fixed
- Default provider path can be overridden via the `SUMMON_PROVIDER_PATH` environment variable,
  resolving an issue where providers cannot be found when installed via homebrew in a non-default location.
  [cyberark/summon#213](https://github.com/cyberark/summon/issues/213)

## [0.8.4] - 2021-05-04

### Added
- Adds apk package to the release artefacts.
  [cyberark/summon#209](https://github.com/cyberark/summon/issues/209)

## [0.8.3] - 2020-09-25

### Added
- Added preliminary support for building Solaris binaries.
  [cyberark/summon#173](https://github.com/cyberark/summon/issues/173)

### Fixed
- Use of a path for a provider via `--provider` CLI flag or `SUMMON_PROVIDER` env
  variable on Windows with `\` as path separators now correctly works.
  [cyberark/summon#167](https://github.com/cyberark/summon/issues/167)
- Fixed handling of errors in the install script.
  [cyberark/summon#171](https://github.com/cyberark/summon/issues/171)

## [0.8.2] - 2020-06-23
### Added
- Summon now supports a `--version-providers` flag to display the versions of installed providers.
  [cyberark/summon#138](https://github.com/cyberark/summon/issues/138)
- Summon supports a `--up` flag that searches for secrets.yml going up, starting from the
  current working directory. This allows the secrets.yml file to be at any directory depth in a
  project, and it is no longer required to be in the current working directory if not specified
  with the `-f` flag.
  [#122](https://github.com/cyberark/summon/issues/122)

## [0.8.1] - 2020-03-02
### Changed
- Added ability to support empty variables [#124](https://github.com/cyberark/summon/issues/124)

### Added
- Added better errors for unknown tags found in the yaml
- Added ability to set a default variable value with `default='<value>'` tag [#38](https://github.com/cyberark/summon/issues/38)

## [0.8.0] - 2019-09-20
### Changed
- To ensure cleanup of files on non-windows platforms we now remain resident
  until the child is killed or it exits [#106](https://github.com/cyberark/summon/issues/106)
- Updated base Golang version to 1.13
- Made Linux builds create static binaries [#65](https://github.com/cyberark/summon/issues/65)

### Added
- Added gitleaks configuration

### Fixed
- Fixed broken website links

## [0.7.0] - 2019-07-11
### Changed
- Updated yaml.v1 dependency to [yaml.v3](gopkg.in/yaml.v3) in part to address
  [cyberark/secretless-broker#785](https://github.com/cyberark/secretless-broker/issues/785)
- Updates goreleaser config to address deprecated sections

### Fixed
- Bumps `ffi` in the `acceptance/` directory to address [this CVE](https://nvd.nist.gov/vuln/detail/CVE-2018-1000201)

### Added
- Added CONTRIBUTING.md for contribution guidelines for the project, including
  contributor agreement

## [0.6.11] - 2019-01-09
### Changed
- Added exporting of `SUMMON_ENV` if `-e` flag is present. Closes[#92](https://github.com/cyberark/summon/issues/92).

## [0.6.10] - 2019-01-03
### Fixed
- Windows subprocess loading is again run with exec.Command. Closes[#88](https://github.com/cyberark/summon/issues/88).

### Changed
- Windows detection of 'Program Files' folder improved.

## [0.6.9] - 2018-12-07
### Changed
- Updated codebase to use [go v1.11 modules](https://github.com/golang/go/wiki/Modules).
- Updated acceptance tests to use an automated test image builds and no makefiles.
- Made subprocess loading take place through execve. Fixes [#83](https://github.com/cyberark/summon/issues/83).

## [0.6.8] - 2018-09-14
### Added
- `-I/--ignore-all` [flag](https://github.com/cyberark/summon#flags), [PR #78](https://github.com/cyberark/summon/pull/78)

## [0.6.7] - 2018-08-06
### Added
- deb and rpm packages

### Changed
- Update build/package process to use [goreleaser](https://github.com/goreleaser/goreleaser).

## [0.6.6] - 2018-02-06
### Changed
- stdout is no longer buffered inside summon. This should greatly decrease the memory footprint of long-running processes wrapped by summon. Closes [#63](https://github.com/cyberark/summon/issues/63).

## [0.6.5] - 2017-08-23
### Changed
- Improved Jenkins CI pipeline.
- Binaries are now built for more distributions (see `PLATFORMS` in [build.sh](build.sh)).
- Simpler docker-compose development environment.

## [0.6.4] - 2017-04-12
### Changed
- Don't rely on executable bit on the provider; instead provide descriptive error if it fails to run - [Issue #40](https://github.com/cyberark/summon/issues/40)

## [0.6.3] - 2017-03-13
### Fixed
- Summon now passes the child exit status to the caller - [PR #39](https://github.com/cyberark/summon/pull/39)

## [0.6.2] - 2017-01-23
### Added
- Added 'default' section support, this is an alias for 'common' - [PR #37](https://github.com/cyberark/summon/pull/37)

## [0.6.1] - 2016-12-19
### Added
- Support Boolean literals - [PR #35](https://github.com/cyberark/summon/pull/35)

## [0.6.0] - 2016-06-20
### Changed
- Write temporary files to home directory if possible

## [0.5.0] - 2016-06-08
### Added
- added `-e`/`--environment` flag

## [0.4.0] - 2016-03-01
### Changed
- **breaking change** Default provider path is now `/usr/local/lib/summon`.

## [0.3.3] - 2016-02-29
### Fixed
- Now fails more gracefully on unknown flags

## [0.3.2] - 2016-02-10
### Fixed
- `@SUMMONENVFILE` is now ensured to contain a trailing newline [GH-22](https://github.com/cyberark/summon/issues/22)

## [0.3.1] - 2016-02-03
### Fixed
- Integer values set in secrets.yml are now parsed correctly [GH-21](https://github.com/cyberark/summon/issues/21)

## [0.3.0] - 2015-08-21
### Changed
- Install bash completions if available
- Switch to tar.gz instead of .zip
- Try to find provider in the default path if just a name given
- Allow -p argument to override SUMMON_PROVIDER envar
- Check if provider exists and is executable

## [0.2.1] - 2015-06-29
### Fixed
- Improve provider path handling [5df0fde](https://github.com/cyberark/summon/commit/5df0fdeb182884371ad647d0a9493a5e07d3e0e4)

## [0.2.0] - 2015-06-23
### Added
- `@SUMMONENVFILE` for better Docker integration

### Fixed
- -D variable interpolation now plays nicely with the shell

## [0.1.2] - 2015-06-10
### Fixed
- Fix --help and --version flags

### Changed
- Vendor dependencies with Godep

## [0.1.1] - 2015-06-09
### Changed
- Attach stdin to allow running interactive processes wrapped with summon
- Changed name from 'cauldron' to 'summon'

## [0.1.0] - 2015-06-03
### Added
- Initial release

[Unreleased]: https://github.com/cyberark/summon/compare/v0.9.7...HEAD
[0.9.7]: https://github.com/cyberark/summon/compare/v0.9.6...v0.9.7
[0.9.6]: https://github.com/cyberark/summon/compare/v0.9.5...v0.9.6
[0.9.5]: https://github.com/cyberark/summon/compare/v0.9.4...v0.9.5
[0.9.4]: https://github.com/cyberark/summon/compare/v0.9.3...v0.9.4
[0.9.3]: https://github.com/cyberark/summon/compare/v0.9.2...v0.9.3
[0.9.2]: https://github.com/cyberark/summon/compare/v0.9.1...v0.9.2
[0.9.1]: https://github.com/cyberark/summon/compare/v0.9.0...v0.9.1
[0.9.0]: https://github.com/cyberark/summon/compare/v0.8.4...v0.9.0
[0.8.4]: https://github.com/cyberark/summon/compare/v0.8.3...v0.8.4
[0.8.3]: https://github.com/cyberark/summon/compare/v0.8.2...v0.8.3
[0.8.2]: https://github.com/cyberark/summon/compare/v0.8.1...v0.8.2
[0.8.1]: https://github.com/cyberark/summon/compare/v0.8.0...v0.8.1
[0.8.0]: https://github.com/cyberark/summon/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/cyberark/summon/compare/v0.6.11...v0.7.0
[0.6.11]: https://github.com/cyberark/summon/compare/v0.6.10...v0.6.11
[0.6.10]: https://github.com/cyberark/summon/compare/v0.6.9...v0.6.10
[0.6.9]: https://github.com/cyberark/summon/compare/v0.6.8...v0.6.9
[0.6.8]: https://github.com/cyberark/summon/compare/v0.6.7...v0.6.8
[0.6.7]: https://github.com/cyberark/summon/compare/v0.6.6...v0.6.7
[0.6.6]: https://github.com/cyberark/summon/compare/v0.6.5...v0.6.6
[0.6.5]: https://github.com/cyberark/summon/compare/v0.6.4...v0.6.5
[0.6.4]: https://github.com/cyberark/summon/compare/v0.6.3...v0.6.4
[0.6.3]: https://github.com/cyberark/summon/compare/v0.6.2...v0.6.3
[0.6.2]: https://github.com/cyberark/summon/compare/v0.6.1...v0.6.2
[0.6.1]: https://github.com/cyberark/summon/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/cyberark/summon/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/cyberark/summon/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/cyberark/summon/compare/v0.3.0...v0.4.0
[0.3.3]: https://github.com/cyberark/summon/compare/v0.3.2...v0.3.3
[0.3.2]: https://github.com/cyberark/summon/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/cyberark/summon/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/cyberark/summon/compare/v0.2.1...v0.3.0
[0.2.1]: https://github.com/cyberark/summon/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/cyberark/summon/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/cyberark/summon/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/cyberark/summon/compare/v0.1.0...v0.1.1
[0.0.1]: https://github.com/cyberark/summon/releases/tag/v0.0.1
