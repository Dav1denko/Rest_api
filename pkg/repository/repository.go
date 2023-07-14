package repository

import (
	restapi "rest_api_TODO"

	"github.com/jmoiron/sqlx"
)

type Authoruzation interface {
	CreateUser(user restapi.User) (int, error)
	GetUser(username, password string) (restapi.User, error)
}

type TodoList interface {
	Create(userId int, list restapi.TodoList) (int, error)
	GetAll(userId int) ([]restapi.TodoList, error)
}

type TodoItem interface{}

type Repository struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authoruzation: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
