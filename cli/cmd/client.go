package cmd

import (
	"github.com/dotindustries/moar/client"
	"github.com/dotindustries/moar/moarpb/v1/v1connect"
)

func protobufClient() v1connect.ModuleRegistryServiceClient {
	return client.New(client.Config{Url: GlobalConfig.BackendAddr})
}
