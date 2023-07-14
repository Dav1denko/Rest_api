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
}

type TodoItem interface{}

type Service struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authoruzation: NewAuthService(repos.Authoruzation),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
