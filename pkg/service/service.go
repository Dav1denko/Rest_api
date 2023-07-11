package service

import (
	restapi "rest_api_TODO"
	"rest_api_TODO/pkg/repository"
)

type Authoruzation interface {
	CreateUser(user restapi.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authoruzation: NewAuthService(repos.Authoruzation),
	}
}
