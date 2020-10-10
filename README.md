# ntpt

Network Time Protocol (NTP) client for testing purposes.

[![Latest Release](https://img.shields.io/github/release/atc0005/ntpt.svg?style=flat-square)](https://github.com/atc0005/ntpt/releases/latest)
[![GoDoc](https://godoc.org/github.com/atc0005/ntpt?status.svg)](https://godoc.org/github.com/atc0005/ntpt)
[![Validate Codebase](https://github.com/atc0005/ntpt/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/ntpt/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/ntpt/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/ntpt/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/ntpt/actions?query=workflow%3A%22Quick+Validation%22)

- [ntpt](#ntpt)
  - [Project home](#project-home)
  - [Overview](#overview)
  - [Changelog](#changelog)
  - [Requirements](#requirements)
  - [How to install it](#how-to-install-it)
  - [Configuration](#configuration)
    - [Command-line arguments](#command-line-arguments)
  - [Examples](#examples)
    - [Basic usage](#basic-usage)
  - [References](#references)

## Project home

See [our GitHub repo](https://github.com/atc0005/ntpt) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

## Overview

Perform a NTP query against a specified server for testing purposes.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

- Go 1.13+ (for building)
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

Tested using:

- Go 1.14+
- Windows 10 Version 1909
  - native
  - WSL
- Ubuntu Linux 18.04+

## How to install it

1. [Download](https://golang.org/dl/) Go
1. [Install](https://golang.org/doc/install) Go
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
