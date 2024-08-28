package app

import (
	"Hexagonal-Architecture-whith-Golang/internal/handler"
	"Hexagonal-Architecture-whith-Golang/internal/infrastructure/router"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	router *fiber.App
}

func NewApp(allhandlers *handler.Handlers) *App {
	app := fiber.New()
	router.SetupRoutes(app, allhandlers)
	return &App{router: app}
}

func (a *App) Run() error {
	return a.router.Listen(":3000")
}
