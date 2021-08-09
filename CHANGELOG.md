# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/ntpt/issues) for any
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

[Unreleased]: https://github.com/atc0005/ntpt/compare/v0.1.3...HEAD
[v0.1.3]: https://github.com/atc0005/ntpt/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/ntpt/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/ntpt/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/ntpt/releases/tag/v0.1.0
