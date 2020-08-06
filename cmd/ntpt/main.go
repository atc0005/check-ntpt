// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/ntpt
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

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
var version string = "dev build"

// Primarily used with branding
const myAppName string = "ntpt"
const myAppURL string = "https://github.com/atc0005/" + myAppName

const (
	// defaultNTPServer is intentionally set to an empty string. Since our
	// intent is to test a specific server, we don't want to hard-code a
	// default NTP server.
	defaultNTPServer string = ""
)

const (
	// ntpServerFlagHelp is the help text provided to the user for that
	// flag option
	ntpServerFlagHelp string = "NTP time server to query"
)

// Branding is responsible for emitting application name, version and origin
func Branding() {
	fmt.Fprintf(
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

		fmt.Fprintf(flag.CommandLine.Output(), "Usage of \"%s\":\n",
			myBinaryName,
		)
		flag.PrintDefaults()

	}
}

func main() {

	var ntpServer string

	flag.StringVar(&ntpServer, "server", defaultNTPServer, ntpServerFlagHelp)
	flag.StringVar(&ntpServer, "s", defaultNTPServer, ntpServerFlagHelp+" (shorthand)")
	flag.Usage = flagsUsage()
	flag.Parse()

	if ntpServer == "" {
		fmt.Println("NTP server not specified!")
		flag.Usage()
		os.Exit(1)
	}

	ntpServerTime, err := ntp.Time(ntpServer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Current time from %s: %+v\n", ntpServer, ntpServerTime)

	// options := ntp.QueryOptions{Timeout: 30 * time.Second, TTL: 5}
	// response, err := ntp.QueryWithOptions(ntpServer, options)
	response, err := ntp.Query(ntpServer)
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"Response from NTP server %q: \n"+
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
		ntpServer,
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
	fmt.Println("Offset adjusted time:", offsetAdjTime)

}
