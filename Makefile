OS = $(shell uname -s)

build:
	@echo "> Building snapshot binary"
	@goreleaser --snapshot --rm-dist
	@echo "> Retrieving the correct binary"
	@mv dist/meereen_${OS}_amd64/* bin/
