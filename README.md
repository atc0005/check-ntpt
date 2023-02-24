# check-ntpt

Go-based tooling to monitor Network Time Protocol (NTP) servers.

[![Latest Release](https://img.shields.io/github/release/atc0005/ntpt.svg?style=flat-square)](https://github.com/atc0005/check-ntpt/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/ntpt.svg)](https://pkg.go.dev/github.com/atc0005/ntpt)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/check-ntpt)](https://github.com/atc0005/check-ntpt)
[![Lint and Build](https://github.com/atc0005/check-ntpt/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/check-ntpt/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/check-ntpt/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/check-ntpt/actions/workflows/project-analysis.yml)

- [check-ntpt](#check-ntpt)
  - [Project home](#project-home)
  - [Overview](#overview)
  - [Changelog](#changelog)
  - [Requirements](#requirements)
    - [Building source code](#building-source-code)
    - [Running](#running)
  - [Installation](#installation)
    - [From source](#from-source)
    - [Using release binaries](#using-release-binaries)
  - [Configuration](#configuration)
    - [Command-line arguments](#command-line-arguments)
  - [Examples](#examples)
    - [Basic usage](#basic-usage)
    - [Verbose output](#verbose-output)
  - [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

Go-based tooling to monitor Network Time Protocol (NTP) servers.

At present, the `ntpt` binary is the only tool provided by this repo. This
tool performs a NTP query against a specified server for testing purposes.

Future plans include providing one (or more) Nagios plugins to monitor NTP
servers.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 10
- Ubuntu Linux 18.04+

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/check-ntpt`
   1. `cd ntpt`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - for current operating system
     - `go build -mod=vendor ./cmd/ntpt/`
       - *forces build to use bundled dependencies in top-level `vendor`
         folder*
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for Windows
      - `make windows`
   - for Linux
     - `make linux`
1. Copy the applicable binary to whatever systems needs to run it
   - if using `Makefile`: look in `/tmp/ntpt/release_assets/ntpt/`
   - if using `go build`: look in `/tmp/ntpt/`

### Using release binaries

1. Download the [latest
   release](https://github.com/atc0005/check-ntpt/releases/latest) binaries
1. Deploy
   - Place `ntpt` in a location of your choice
     - e.g., `/usr/local/bin/`

## Configuration

### Command-line arguments

- Flags marked as **`required`** must be set via CLI flag *or* within a
  TOML-formatted configuration file.
- Flags *not* marked as required are for settings where a useful default is
  already defined.

| Flag           | Required | Default        | Repeat | Possible                                              | Description                                             |
| -------------- | -------- | -------------- | ------ | ----------------------------------------------------- | ------------------------------------------------------- |
| `h`, `help`    | No       | `false`        | No     | `h`, `help`                                           | Show Help text along with the list of supported flags.  |
| `s`, `server`  | **Yes**  | *empty string* | **No** | *one valid IP Address or fully-qualified server name* | NTP server to submit query against.                     |
| `v`, `verbose` | No       | `false`        | **No** | `true`, `false`                                       | Enables display of verbose output. Disabled by default. |

## Examples

### Basic usage

```console
$ ./ntpt -s pool.ntp.org
Current time from pool.ntp.org: 2022-08-15 06:16:11.139364388 -0500 CDT m=-0.121229211
Current time from local system: 2022-08-15 06:16:11.5944353 -0500 CDT m=+0.333841701

The local system is -456.993422ms behind pool.ntp.org.
```

### Verbose output

```console
$ ./ntpt -v -s pool.ntp.org
Current time from pool.ntp.org: 2022-08-15 06:16:41.864305139 -0500 CDT m=-0.193111660
Current time from local system: 2022-08-15 06:16:42.3228386 -0500 CDT m=+0.265421801

Response from NTP server "pool.ntp.org":
        Time: 2022-08-15 11:16:41.968756524 +0000 UTC
        ClockOffset: -459.00146ms
        RTT: 203.878832ms
        Stratum: 2
        ReferenceID: 167864580
        ReferenceTime: 2022-08-15 11:16:39.020873716 +0000 UTC
        RootDelay: 213.623µs
        RootDispersion: 30.518µs
        RootDistance: 102.076745ms
        Leap: 0
        MinError: 357.062044ms
        KissCode: ""

Offset adjusted time: 2022-08-15 06:16:42.07080214 -0500 CDT m=+0.013385341
```

## References

- <https://github.com/beevik/ntp>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/check-ntpt>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
