package persistence

import (
	domain "Hexagonal-Architecture-whith-Golang/internal/domain/entities"
	"Hexagonal-Architecture-whith-Golang/internal/repository"
	"errors"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) repository.UserRepository {
	return &PostgresUserRepository{db}
}

func (r *PostgresUserRepository) GetUserByID(id int) (*domain.User, error) {
	if id == 1 {
		return &domain.User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}, nil
	}
	return nil, errors.New("user not found")
}
