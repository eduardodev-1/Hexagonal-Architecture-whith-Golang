package service

import (
	domain "Hexagonal-Architecture-whith-Golang/internal/domain/entities"
	"Hexagonal-Architecture-whith-Golang/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}
