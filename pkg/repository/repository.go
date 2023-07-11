package repository

import (
	restapi "rest_api_TODO"

	"github.com/jmoiron/sqlx"
)

type Authoruzation interface {
	CreateUser(user restapi.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authoruzation: NewAuthPostgres(db),
	}
}
