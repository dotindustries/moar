protoc:
	protoc --twirp_out=./moarpb --go_out=./moarpb --go_opt=paths=source_relative --twirp_opt=paths=source_relative moar.proto

swagger:
	twirp-swagger-gen -in moar.proto -out docs/moar.swagger.json -host localhost:8000

build:
	go build -o moar cli/main.go
