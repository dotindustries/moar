protoc:
	protoc --twirp_out=./moarpb --go_out=./moarpb --go_opt=paths=source_relative --twirp_opt=paths=source_relative moar.proto

swagger: install-deps
	twirp-swagger-gen -in moar.proto -out docs/moar.swagger.json -host localhost:8000

install-deps:
	go install github.com/go-bridget/twirp-swagger-gen/cmd/twirp-swagger-gen@latest

build:
	go build -o moar cli/main.go

all: protoc swagger build