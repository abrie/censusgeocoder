GOPATH=$(shell pwd)
GOBIN=$(shell pwd)/bin

build:
	@GOPATH=$(GOPATH) go build -o bin/cli github.com/abrie/censusgeocoder/cmd/cli
