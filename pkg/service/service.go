package service

import "rest_api_TODO/pkg/repository"

type Authoruzation interface{}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
