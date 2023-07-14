package service

import (
	restapi "rest_api_TODO"
	"rest_api_TODO/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list restapi.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
