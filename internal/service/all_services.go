package service

import (
	"Hexagonal-Architecture-whith-Golang/internal/repository"
)

type Services struct {
	UserService *UserService
}

func NewServices(allRepositories *repository.Repositories) *Services {
	return &Services{
		UserService: NewUserService(allRepositories.UserRepository),
	}
}
