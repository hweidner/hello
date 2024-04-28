# Copyright 2019-2020, Harald Weidner and the hello contributors
# SPDX-License-Identifier: MIT

VERSION = $(shell git describe --always --dirty --tags)
GOFLAGS = -trimpath
LDFLAGS = -X main.Version=$(VERSION)

hello: hello.go
	go build $(GOFLAGS) -ldflags "$(LDFLAGS)"

run: hello
	./hello

clean:
	rm -f hello *~

.PHONY: run clean
