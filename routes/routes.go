package routes

import (
	"todo/db"
	"todo/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	r := e.Group("/")

	r.Any("", handlers.PlaygroundHandler())
	r.Any("ws", handlers.PlaygroundWsHandler())

	r.Any("query", handlers.GraphqlHandler(db.DB))
	r.Any("subscriptions", handlers.GraphqlWsHandler(db.DB))
}
