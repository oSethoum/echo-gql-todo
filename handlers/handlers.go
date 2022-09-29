package handlers

import (
	"net/http"
	"time"
	"todo/ent"
	"todo/graph/resolvers"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func PlaygroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func PlaygroundWsHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL WS", "/subscription")
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func GraphqlHandler(client *ent.Client) echo.HandlerFunc {
	h := handler.NewDefaultServer(resolvers.ExecutableSchema())
	h.Use(entgql.Transactioner{TxOpener: client})
	h.Use(extension.Introspection{})
	h.AddTransport(transport.POST{})
	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
