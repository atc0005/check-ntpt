// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-ntpt
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/beevik/ntp"
)

// Overridden via Makefile for release builds
var version = "dev build"

// Primarily used with branding
const myAppName string = "ntpt"
const myAppURL string = "https://github.com/atc0005/" + myAppName

// Default flag settings if not overridden by user input.
const (
	// defaultNTPServer is intentionally set to an empty string. Since our
	// intent is to test a specific server, we don't want to hard-code a
	// default NTP server.
	defaultNTPServer string = ""

	defaultVerboseOutput bool = false
)

// Flag help text (user visible).
const (
	ntpServerFlagHelp     string = "NTP time server to query"
	verboseOutputFlagHelp string = "Enables display of verbose output. Disabled by default."
)

type config struct {
	server  string
	verbose bool
}

// Branding is responsible for emitting application name, version and origin
func Branding() {
	_, _ = fmt.Fprintf(
		flag.CommandLine.Output(),
		"\n%s %s\n%s\n\n",
		myAppName,
		version,
		myAppURL,
	)
}

// flagsUsage displays branding information and general usage details
func flagsUsage() func() {

	return func() {

		myBinaryName := filepath.Base(os.Args[0])

		Branding()

		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of \"%s\":\n",
			myBinaryName,
		)
		flag.PrintDefaults()

	}
}

func main() {

	var cfg config

	flag.StringVar(&cfg.server, "server", defaultNTPServer, ntpServerFlagHelp)
	flag.StringVar(&cfg.server, "s", defaultNTPServer, ntpServerFlagHelp+" (shorthand)")
	flag.BoolVar(&cfg.verbose, "verbose", defaultVerboseOutput, verboseOutputFlagHelp)
	flag.BoolVar(&cfg.verbose, "v", defaultVerboseOutput, verboseOutputFlagHelp+" (shorthand)")
	flag.Usage = flagsUsage()
	flag.Parse()

	if cfg.server == "" {
		fmt.Println("NTP server not specified!")
		flag.Usage()
		os.Exit(1)
	}

	ntpServerTime, err := ntp.Time(cfg.server)
	if err != nil {
		panic(err)
	}

	localTime := time.Now()

	fmt.Printf("Current time from %s: %v\n", cfg.server, ntpServerTime)
	fmt.Printf("Current time from local system: %v\n", localTime)

	// options := ntp.QueryOptions{Timeout: 30 * time.Second, TTL: 5}
	// response, err := ntp.QueryWithOptions(ntpServer, options)
	response, err := ntp.Query(cfg.server)
	if err != nil {
		panic(err)
	}

	// fmt.Printf(
	// 	"Estimated offset of local client clock relative to %s: %+v\n",
	// 	cfg.server,
	// 	response.ClockOffset,
	// )

	switch {
	case cfg.verbose:
		fmt.Printf(
			"\nResponse from NTP server %q: \n"+
				"\tTime: %v\n"+
				"\tClockOffset: %v\n"+
				"\tRTT: %v\n"+
				"\tStratum: %v\n"+
				"\tReferenceID: %v\n"+
				"\tReferenceTime: %v\n"+
				"\tRootDelay: %v\n"+
				"\tRootDispersion: %v\n"+
				"\tRootDistance: %v\n"+
				"\tLeap: %v\n"+
				"\tMinError: %v\n"+
				"\tKissCode: %q\n",
			cfg.server,
			response.Time,
			response.ClockOffset,
			response.RTT,
			response.Stratum,
			response.ReferenceID,
			response.ReferenceTime,
			response.RootDelay,
			response.RootDispersion,
			response.RootDistance,
			response.Leap,
			response.MinError,
			response.KissCode,
		)

		offsetAdjTime := time.Now().Add(response.ClockOffset)
		fmt.Printf("\nOffset adjusted time: %v\n", offsetAdjTime)

	default:

		timeDirection := "behind"
		if localTime.Before(ntpServerTime) {
			timeDirection = "ahead of"
		}

		// The local system is 2m28.258534398s behind pool.ntp.org.
		// The local system is 2m28.258534398s ahead of pool.ntp.org.
		fmt.Printf(
			"\nThe local system is %v %s %s.\n",
			response.ClockOffset,
			timeDirection,
			cfg.server,
		)

	}

}
