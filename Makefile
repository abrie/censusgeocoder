GOPATH=$(shell pwd)
GOBIN=$(shell pwd)/bin

dependencies:
	@GOPATH=$(GOPATH) go get -v -t -d ./...

build:
	@GOPATH=$(GOPATH) go build -v -o bin/cli github.com/abrie/censusgeocoder/cmd/cli

test:
	@GOPATH=$(GOPATH) go test -v ./...
