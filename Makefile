#!/bin/make
GOROOT:=$(shell PATH="/pkg/main/dev-lang.go.dev/bin:$$PATH" go env GOROOT)
GOPATH:=$(shell $(GOROOT)/bin/go env GOPATH)

.PHONY: test deps

all:
	GOROOT="$(GOROOT)" $(GOPATH)/bin/goimports -w -l .
	$(GOROOT)/bin/go build -v

deps:
	$(GOROOT)/bin/go get -v -t .

bench:
	$(GOROOT)/bin/go test -tags bench -benchmem -bench .

test:
	$(GOROOT)/bin/go test -v ./...
