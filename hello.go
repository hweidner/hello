// Copyright 2019-2020, Harald Weidner and the hello contributors
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
	"runtime/debug"
)

var (
	// Version contains the repository version, as reported by: git describe --always --dirty --tags
	Version = ""
)

func main() {
	fmt.Println("Hello, Gophers!")
	fmt.Println()
	if Version != "" {
		fmt.Println("Git version:", Version)
	}
	runtimeInfo := getVersion()
	fmt.Println(runtimeInfo)
}

type versionInfo struct {
	goVersion            string
	goos, goarch         string
	gomaxprocs           int
	vcsrevision, vcstime string
}

func getVersion() *versionInfo {
	revision, time := vcsRevision()
	return &versionInfo{
		goVersion:   runtime.Version(),
		goos:        runtime.GOOS,
		goarch:      runtime.GOARCH,
		gomaxprocs:  runtime.GOMAXPROCS(0),
		vcsrevision: revision,
		vcstime:     time,
	}

}

func (v *versionInfo) String() string {
	format := `Go version:  %s
Platform:    %s/%s
GOMAXPROCS:  %d
vcsRevision: %s
vcsTime:     %s
`

	return fmt.Sprintf(format, v.goVersion, v.goos, v.goarch, v.gomaxprocs, v.vcsrevision, v.vcstime)
}

func vcsRevision() (revision, time string) {
	modified := ""

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				revision = setting.Value
			case "vcs.time":
				time = setting.Value
			case "vcs.modified":
				modified = setting.Value
			}
		}
	}
	if modified == "true" {
		time = time + " (modified)"
	}
	return
}
