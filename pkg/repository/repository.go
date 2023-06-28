package repository

type Authoruzation interface{}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authoruzation
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
