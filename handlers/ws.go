package handlers

import (
	"net/http"
	"time"
	"todo/ent"
	"todo/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func GraphqlWs_Handler(client *ent.Client) echo.HandlerFunc {
	h := handler.New(resolvers.ExecutableSchema())
	h.Use(extension.Introspection{})
	h.AddTransport(transport.POST{})
	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		// InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
		// 	return auth.WebSocketInit(ctx, initPayload)
		// },
	})
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
