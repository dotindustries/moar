protoc:
	buf generate

# TODO: replace this with the appropriate buf/connect/swagger variant?
swagger: install-deps
	twirp-swagger-gen -in moar.proto -out docs/moar.swagger.json -host localhost:8000

install-deps:
	go install github.com/sudorandom/protoc-gen-connect-openapi@main

VERSION=$(shell git describe --abbrev=0 --tags)
COMMIT=$(shell git rev-parse --short HEAD)

build:
	go build -o moar "-ldflags=-X main.Version=$(VERSION) -X main.CommitHash=$(COMMIT)" cli/main.go

all: protoc swagger build