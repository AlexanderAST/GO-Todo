package service

import (
	"todo"
	"todo/pkg/repository"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(todoL todo.Todo) (int, error) {
	return s.repo.CreateTodo(todoL)
}

func (s *TodoService) AllContent() ([]*todo.Todo, error) {
	code, err := s.repo.AllContent()
	return code, err
}

func (s *TodoService) DeleteTodo(title string) (string, error) {
	str, _ := s.repo.DeleteTodo(title)
	return str, nil
}

func (s *TodoService) UpdateTodo(todo todo.Todo) (string, error) {
	str, _ := s.repo.UpdateTodo(todo)
	return str, nil
}
