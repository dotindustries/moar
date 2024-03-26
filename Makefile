protoc:
	buf generate

# TODO: replace this with the appropriate buf/connect/swagger variant?
swagger: install-deps
	twirp-swagger-gen -in moar.proto -out docs/moar.swagger.json -host localhost:8000

install-deps:


build:
	go build -o moar cli/main.go

all: protoc swagger build