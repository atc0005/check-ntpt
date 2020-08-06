/*

Network Time Protocol (NTP) client for testing purposes.

PROJECT HOME

See our GitHub repo (https://github.com/atc0005/ntpt) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

PURPOSE

Perform a NTP query against a specified server for testing purposes.

FEATURES

• single binary, no outside dependencies

• query specified NTP server for current time

USAGE

	$ ./ntpt -h

	ntpt dev build
	https://github.com/atc0005/ntpt

	Usage of "ntpt":
	-s string
			NTP time server to query (shorthand)
	-server string
			NTP time server to query

EXAMPLE

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

*/
package main
