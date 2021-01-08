// Copyright 2019 - 2020, Harald Weidner and the hello contributors
// SPDX-License-Identifier: MIT

/*
Command hello greets you with a friendly message.

Command hello prints a friendly greeting message, and some informations about the
Go runtime.
This repository is just there to test the new Go developer site https://go.dev/
*/
package main

import (
	"fmt"
	"runtime"
)

var (
	// Version contains the repository version, as reported by: git describe --always --dirty --tags
	Version = "(undefined)"
	// BuildDate is the date of the build in RFC3339 format.
	BuildDate = "(undefined)"
)

func main() {
	fmt.Println("Hello, Gophers!\n")
	fmt.Println("Hello version:", Version)
	fmt.Println("Build date:", BuildDate, "\n")

	version := getVersion()
	fmt.Println("Runtime information:")
	fmt.Println(version)
}

type versionInfo struct {
	goVersion    string
	goos, goarch string
	gomaxprocs   int
}

func getVersion() *versionInfo {
	return &versionInfo{
		goVersion:  runtime.Version(),
		goos:       runtime.GOOS,
		goarch:     runtime.GOARCH,
		gomaxprocs: runtime.GOMAXPROCS(0),
	}
}

func (v *versionInfo) String() string {
	return fmt.Sprintf("Go version: %s\nPlatform: %s/%s\nGOMAXPROCS: %d\n",
		v.goVersion, v.goos, v.goarch, v.gomaxprocs)
}
