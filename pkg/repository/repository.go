package repository

import (
	"github.com/jmoiron/sqlx"
	"todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	AuthUser(username, password string) (todo.User, error)
}

type Todo interface {
	CreateTodo(todoL todo.Todo) (int, error)
	AllContent() ([]*todo.Todo, error)
	DeleteTodo(title string) (string, error)
	UpdateTodo(todo todo.Todo) (string, error)
}

type Repository struct {
	Authorization
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Todo:          NewTodoPostgres(db),
	}
}
