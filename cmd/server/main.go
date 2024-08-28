package main

import (
	application "Hexagonal-Architecture-whith-Golang/internal/app"
	"Hexagonal-Architecture-whith-Golang/internal/handler"
	"Hexagonal-Architecture-whith-Golang/internal/infrastructure/database"
	"Hexagonal-Architecture-whith-Golang/internal/repository"
	"Hexagonal-Architecture-whith-Golang/internal/service"
	"log"
)

func main() {
	db := database.NewPosgreSQLConnection()

	allRepositories := repository.NewRepositories(db, repository.Postgresql)

	allServices := service.NewServices(allRepositories)

	allHandlers := handler.NewHandlers(allServices)

	app := application.NewApp(allHandlers)
	if err := app.Run(); err != nil {
		log.Fatal("Failed to start the app:", err)
	}
}
