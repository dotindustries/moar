package client

import (
	"connectrpc.com/connect"
	"github.com/dotindustries/moar/moarpb/v1/moarpbconnect"
	"net/http"
)

type Config struct {
	Url        string
	HttpClient *http.Client
}

// New creates a new protobuf client. If no http client is specified it will fall back to the default http.Client
func New(config Config, opts ...connect.ClientOption) moarpbconnect.ModuleRegistryServiceClient {
	httpCli := config.HttpClient
	if httpCli == nil {
		httpCli = http.DefaultClient
	}
	return moarpbconnect.NewModuleRegistryServiceClient(
		httpCli,
		config.Url,
	)
}
