#!/bin/make
GOPATH:=$(shell go env GOPATH)

.PHONY: test bench

all:
	$(GOPATH)/bin/goimports -w -l .
	go build -v

test:
	go test ./...

bench:
	go test -tags bench -benchmem -bench .

install:
	go get ./...
