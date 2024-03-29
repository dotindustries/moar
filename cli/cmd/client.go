package cmd

import (
	"connectrpc.com/connect"
	"context"
	"github.com/dotindustries/moar/client"
	"github.com/dotindustries/moar/moarpb/v1/v1connect"
	"github.com/labstack/echo/v4"
)

var clientAuthProviderInterceptor = connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		if request.Spec().IsClient && GlobalConfig.AccessToken != "" {
			// Send a token with client requests.
			request.Header().Set(echo.HeaderAuthorization, GlobalConfig.AccessToken)
		}
		return next(ctx, request)
	}
})

func protobufClient() v1connect.ModuleRegistryServiceClient {
	return client.New(client.Config{Url: GlobalConfig.BackendAddr}, connect.WithInterceptors(clientAuthProviderInterceptor))
}
