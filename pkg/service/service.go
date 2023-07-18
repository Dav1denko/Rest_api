package service

import (
	restapi "rest_api_TODO"
	"rest_api_TODO/pkg/repository"
)

type Authoruzation interface {
	CreateUser(user restapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list restapi.TodoList) (int, error)
	GetAll(userId int) ([]restapi.TodoList, error)
	GetById(userId, listId int) (restapi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input restapi.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item restapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]restapi.TodoItem, error)
	GetById(userId, itemId int) (restapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input restapi.UpdateItemInput) error
}

type Service struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authoruzation: NewAuthService(repos.Authoruzation),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
