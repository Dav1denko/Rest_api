package repository

import "github.com/jmoiron/sqlx"

type Authoruzation interface{}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
