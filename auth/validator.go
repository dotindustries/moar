package auth

import (
	"errors"
	unkey "github.com/WilfredAlmeida/unkey-go/features"
	"github.com/labstack/echo/v4"
)

const ContextKey = "auth"

func KeyValidator(key string, c echo.Context) (bool, error) {
	if key == "" {
		return false, errors.New("missing API key")
	}
	resp, err := unkey.KeyVerify(key)
	if err != nil {
		return false, err
	}
	if !resp.Valid {
		return false, nil
	}

	c.Set(ContextKey, resp)
	return true, nil
}

func FromContext(c echo.Context) unkey.KeyVerifyResponse {
	return c.Get(ContextKey).(unkey.KeyVerifyResponse)
}
