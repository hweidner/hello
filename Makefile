# Copyright 2019-2020, Harald Weidner and the hello contributors
# SPDX-License-Identifier: MIT

VERSION = $(shell git describe --always --dirty --tags)
DATE = $(shell date -Iseconds)
LDFLAGS = "-X main.Version=$(VERSION) -X main.BuildDate=$(DATE)"

hello: hello.go
	go build -ldflags $(LDFLAGS)

test: hello
	./hello

clean:
	rm -f hello *~

git:
	git add .
	git commit
	git push

.PHONY: test clean git
