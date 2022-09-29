package routes

import (
	"todo/auth"
	"todo/db"
	"todo/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.POST("/login", auth.Login)

	r := e.Group("/")
	r.Use(auth.ValuesToContext)

	r.Any("", handlers.PlaygroundHandler())
	r.Any("ws", handlers.PlaygroundWsHandler())

	r.Any("query", handlers.GraphqlHandler(db.DB))
	r.Any("subscriptions", handlers.GraphqlHandler(db.DB))
}
