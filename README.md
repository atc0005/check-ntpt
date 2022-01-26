# ntpt

Network Time Protocol (NTP) client for testing purposes.

[![Latest Release](https://img.shields.io/github/release/atc0005/ntpt.svg?style=flat-square)](https://github.com/atc0005/ntpt/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/ntpt.svg)](https://pkg.go.dev/github.com/atc0005/ntpt)
[![Validate Codebase](https://github.com/atc0005/ntpt/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/ntpt/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/ntpt/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/ntpt/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Quick+Validation%22)

- [ntpt](#ntpt)
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
  - [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

Perform a NTP query against a specified server for testing purposes.

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
   1. `git clone https://github.com/atc0005/ntpt`
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
   release](https://github.com/atc0005/ntpt/releases/latest) binaries
1. Deploy
   - Place `ntpt` in a location of your choice
     - e.g., `/usr/local/bin/`

## Configuration

### Command-line arguments

- Flags marked as **`required`** must be set via CLI flag *or* within a
  TOML-formatted configuration file.
- Flags *not* marked as required are for settings where a useful default is
  already defined.

| Flag          | Required | Default        | Repeat | Possible                                              | Description                                            |
| ------------- | -------- | -------------- | ------ | ----------------------------------------------------- | ------------------------------------------------------ |
| `h`, `help`   | No       | `false`        | No     | `h`, `help`                                           | Show Help text along with the list of supported flags. |
| `s`, `server` | **Yes**  | *empty string* | **No** | *one valid IP Address or fully-qualified server name* | NTP server to submit query against.                    |

## Examples

### Basic usage

```ShellSession
$ ./ntpt --server ntp.example.com
Current time from ntp.example.com: 2020-08-06 04:28:26.820043274 -0500 CDT m=+0.183802875
Response from NTP server "ntp.example.com":
        Time: 2020-08-06 09:28:26.905991848 +0000 UTC
        ClockOffset: 31.803212ms
        RTT: 135.462329ms
        Stratum: 1
        ReferenceID: 1196446464
        ReferenceTime: 2020-08-06 09:28:26.418643103 +0000 UTC
        RootDelay: 0s
        RootDispersion: 991.821µs
        RootDistance: 68.722985ms
        Leap: 0
        MinError: 0s
        KissCode: ""
Offset adjusted time: 2020-08-06 04:28:26.977070412 -0500 CDT m=+0.340829713
```

## References

- <https://github.com/beevik/ntp>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/ntpt>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
