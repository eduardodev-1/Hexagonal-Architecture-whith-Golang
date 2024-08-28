package repository

import (
	"Hexagonal-Architecture-whith-Golang/internal/infrastructure/persistence/postgres"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	Postgresql = "postgresql"
)

type Repositories struct {
	UserRepository UserRepository
}

func NewRepositories(db interface{}, dbType string) *Repositories {
	switch dbType {
	case Postgresql:
		return &Repositories{
			UserRepository: persistence.NewPostgresUserRepository(db.(*sqlx.DB)),
		}
	default:
		log.Fatal("unsupported db type")
	}
	return nil
}
