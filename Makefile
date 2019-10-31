VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard cmd/*.go)

PROJECTNAME := $(shell basename "$(GOBASE)")

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

build:
	@echo "> Building binary"
	@GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

clean:
	@echo "> Cleaning build cache"
	@rm $(GOBIN)/$(PROJECTNAME)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
