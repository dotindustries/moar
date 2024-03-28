package auth

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	unkey "github.com/WilfredAlmeida/unkey-go/features"
	"github.com/labstack/echo/v4"
	"strings"
)

const ContextKey = "auth"
const bearerAuthScheme = "Bearer "

var ApiKeyInterceptor = connect.UnaryInterceptorFunc(
	func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			key := req.Header().Get(echo.HeaderAuthorization)
			if key == "" {
				key = req.Header().Get("X-Api-Key")
			}
			key = strings.TrimPrefix(key, bearerAuthScheme)
			if key == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("missing API key"))
			}

			resp, err := unkey.KeyVerify(key)
			if err != nil {
				return nil, err
			}
			if !resp.Valid {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid API key"))
			}

			// set auth context key
			ctx = context.WithValue(ctx, ContextKey, resp)
			res, err := next(ctx, req)
			return res, err
		})
	},
)

func FromContext(ctx context.Context) unkey.KeyVerifyResponse {
	return ctx.Value(ContextKey).(unkey.KeyVerifyResponse)
}
