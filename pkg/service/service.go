package service

import (
	"todo"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (todo.User, error)
}

type Todo interface {
	CreateTodo(todoL todo.Todo) (int, error)
	AllContent() ([]*todo.Todo, error)
	DeleteTodo(title string) (string, error)
	UpdateTodo(todo todo.Todo) (string, error)
}

type Service struct {
	Authorization
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Todo:          NewTodoService(repos.Todo),
	}
}
