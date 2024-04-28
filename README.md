[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDocs](https://godocs.io/github.com/hweidner/hello?status.svg)](https://godocs.io/github.com/hweidner/hello)
[![Go Reference](https://pkg.go.dev/badge/github.com/hweidner/hello.svg)](https://pkg.go.dev/github.com/hweidner/hello)
[![Go Report Card](https://goreportcard.com/badge/github.com/hweidner/hello)](https://goreportcard.com/report/github.com/hweidner/hello)

# hello

Command hello greets you with a friendly message.

This repository is just there to test the new Go developer site [go.dev](https://go.dev/).

## Installation

You need a working Go environment on you system. Use ``sudo apt install golang``
if you run a Debian or Ubuntu Linux system. Then install this repository by

	$ go get github.com/hweidner/hello

This command will install the program to ``~/go/bin/hello``, or
``$GOPATH/bin/hello`` if ``GOPATH`` is set. Then run it:

Alernatively, pull the repository and execute `make`.

When executed, the program prints a message like this. Note that the `Git version` line is only
printed if the program was compiled using `make`.

    $ ./hello
    Hello, Gophers!

    Git version: v0.5.0-1-ga451f73-dirty
    Go version:  go1.22.2
    Platform:    linux/amd64
    GOMAXPROCS:  16
    vcsRevision: a451f735655f4507ceea6beed1dd3ae594e4f930
    vcsTime:     2024-04-28T01:24:48Z (modified)

## License

This software is released under the MIT license. See the file [LICENSE](LICENSE)
for a full text of the MIT license.
