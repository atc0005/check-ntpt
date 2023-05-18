# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/check-ntpt/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.2.1] - 2023-05-18

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.9`
  - `beevik/ntp`
    - `v0.3.0` to `v1.0.0`
  - `golang.org/x/net`
    - `v0.4.0` to `v0.10.0`
  - `golang.org/x/sys`
    - `v0.3.0` to `v0.8.0`
- (GH-110) Add Go Module Validation, Dependency Updates jobs
- (GH-115) Drop `Push Validation` workflow
- (GH-116) Rework workflow scheduling
- (GH-118) Remove `Push Validation` workflow status badge

### Fixed

- (GH-124) Update vuln analysis GHAW to use on.push hook

## [v0.2.0] - 2022-12-12

### Overview

- Add new flag
- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-84) Simplify default output, add verbose flag to enable emitting current
  output

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.19.4`
  - `golang.org/x/net`
    - `v0.0.0-20200707034311-ab3426394381` to `v0.4.0`
  - `golang.org/x/sys`
    - `v0.0.0-20210927094055-39ccf1dd6fa6` to `v0.3.0`
- (GH-88) Update project to Go 1.19
- (GH-90) Update .gitignore exclusions
- (GH-91) Update Makefile and GitHub Actions Workflows
- (GH-93) Rename project from ntpt to check-ntpt
- (GH-98) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-83) Update lintinstall Makefile recipe
- (GH-89) Add missing cmd doc file
- (GH-102) Fix verbose flag help output
- (GH-103) Fix Makefile Go module base path detection

## [v0.1.7] - 2022-03-03

### Overview

- Dependency updates
- CI / linting improvements
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-68) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-69) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

### Fixed

- (GH-71) var-declaration: should omit type string from declaration of var
  (revive)

## [v0.1.6] - 2022-01-26

### Overview

- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - (GH-63) Update go.mod file, canary Dockerfile to reflect current
      dependencies

## [v0.1.5] - 2021-12-29

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.1`

## [v0.1.4] - 2021-11-10

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.10`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

- (GH-49) Lock Go version to the latest "oldstable" series

## [v0.1.3] - 2021-08-09

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version
    - updated from `v2.2.0` to `v2.4.0`

## [v0.1.2] - 2021-07-19

### Overview

- Dependency updates
- Minor fixes
- Built using Go 1.16.6
  - **Statically linked**
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- Add "canary" Dockerfile to track stable Go releases, serve as a reminder to
  generate fresh binaries

### Changed

- Swap out GoDoc badge for pkg.go.dev badge

- Dependencies
  - `Go`
    - `1.15.2` to `1.16.6`
  - `actions/checkout`
    - `v2.3.3` to `v2.3.4`
  - `actions/setup-node`
    - `v2.1.2` to `v2.2.0`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

## [v0.1.1] - 2020-10-11

### Added

- Binary release
  - Built using Go 1.15.2
  - **Statically linked**
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

### Changed

- Dependencies
  - `actions/checkout`
    - `v2.3.1` to `v2.3.3`
  - `actions/setup-node`
    - `v2.1.1` to `v2.1.2`

- Add `-trimpath` build flag

### Fixed

- Makefile build options do not generate static binaries
- Makefile generates checksums with qualified path
- Typo in README
- Missing doc comment

## [v0.1.0] - 2020-08-06

### Added

Initial release!

- Command-line flags support via standard library `flag` package
- Go modules (vs classic GOPATH setup)
- Vendored dependencies

- README
  - Link badges to applicable GitHub Actions workflows results

- Add Docker-based GitHub Actions Workflows
  - Use containers created and managed through the `atc0005/go-ci`
    project.

  - Primary workflow
    - with parallel linting, testing and building tasks
    - with three Go environments
      - "old stable" - currently `Go 1.13.14`
      - "stable" - currently `Go 1.14.6`
      - "unstable" - currently `Go 1.15rc1`
    - Makefile is *not* used in this workflow
    - staticcheck linting using latest stable version provided by the
      `atc0005/go-ci` containers

  - Separate Makefile-based linting and building workflow
    - intended to help ensure that local Makefile-based builds that
      are referenced in project README files continue to work as
      advertised until a better local tool can be discovered/explored
      further
    - use `golang:latest` container to allow for Makefile-based
      linting tooling installation testing since the `atc0005/go-ci`
      project provides containers with those tools already
      pre-installed
      - linting tasks use container-provided `golangci-lint` config
        file *except* for the Makefile-driven linting task which
        continues to use the repo-provided copy of the `golangci-lint`
        configuration file

  - Add Quick Validation workflow
    - run on every push, everything else on pull request updates
    - linting via `golangci-lint` only
    - testing
    - no builds

- Dependabot updates

- Makefile driven builds and linting

[Unreleased]: https://github.com/atc0005/check-ntpt/compare/v0.2.1...HEAD
[v0.2.1]: https://github.com/atc0005/check-ntpt/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/check-ntpt/releases/tag/v0.2.0
[v0.1.7]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.7
[v0.1.6]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.6
[v0.1.5]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.5
[v0.1.4]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.4
[v0.1.3]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/check-ntpt/releases/tag/v0.1.0
