package repository

import domain "Hexagonal-Architecture-whith-Golang/internal/domain/entities"

type UserRepository interface {
	GetUserByID(id int) (*domain.User, error)
}
