package client

import (
	"net/http"

	"github.com/dotindustries/moar/moarpb"
	"github.com/twitchtv/twirp"
)

type Config struct {
	Url        string
	HttpClient *http.Client
}

// New creates a new protobuf client. If no http client is specified it will fall back to the default http.Client
func New(config Config, opts ...twirp.ClientOption) moarpb.ModuleRegistry {
	httpCli := config.HttpClient
	if httpCli == nil {
		httpCli = &http.Client{}
	}
	return moarpb.NewModuleRegistryProtobufClient(config.Url, httpCli, opts...)
}

// NewJSON creates a new JSON client. If no http client is specified it will fall back to the default http.Client
func NewJSON(config Config, opts ...twirp.ClientOption) moarpb.ModuleRegistry {
	httpCli := config.HttpClient
	if httpCli == nil {
		httpCli = &http.Client{}
	}
	return moarpb.NewModuleRegistryJSONClient(config.Url, httpCli, opts...)
}
