package handler

import (
	"Hexagonal-Architecture-whith-Golang/internal/service"
)

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(allServices *service.Services) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(allServices.UserService),
	}
}
