package auth

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Protected() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &Claims{},
		SigningKey: []byte("secret"),
	}

	return middleware.JWTWithConfig(config)
}

func ValuesToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// extracting default context from echo
		ctx := c.Request().Context()

		if c.IsWebSocket() {
			// we usually get the socket client id in case of anonymuous chat ( doesn't require authentication )
			ctx = context.WithValue(ctx, ContextKey{Key: "socketClient"}, c.Request().Header.Get("Sec-Websocket-Key"))
		}

		c.SetRequest(c.Request().WithContext(ctx))

		next(c)
		return nil
	}
}
