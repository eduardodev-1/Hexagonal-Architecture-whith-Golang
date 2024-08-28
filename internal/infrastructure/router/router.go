package router

import (
	"Hexagonal-Architecture-whith-Golang/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handlers *handler.Handlers) {
	v1 := app.Group("/v1")
	user := v1.Group("/user")
	{
		user.Get("/:id", handlers.UserHandler.GetUser)
	}
}
