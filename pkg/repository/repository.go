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
	GetById(userId, listId int) (restapi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input restapi.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item restapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]restapi.TodoItem, error)
	GetById(userId, itemId int) (restapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input restapi.UpdateItemInput) error
}

type Repository struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authoruzation: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
