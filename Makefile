#!/bin/make
GOPATH:=$(shell go env GOPATH)

.PHONY: deps testdeps test bench

all:
	$(GOPATH)/bin/goimports -w -l .
	go build -v

test:
	go test ./...

bench:
	go test -tags bench -benchmem -bench .

deps:
	go get ./...

testdeps:
	go get -t ./...
