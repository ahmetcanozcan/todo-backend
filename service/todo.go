package service

import (
	"app/model"
	"app/repository"
	"context"
)

type TodoService interface {
	GetTodos(context.Context) ([]model.Todo, error)
	AddTodo(context.Context, model.Todo) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) GetTodos(ctx context.Context) ([]model.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) AddTodo(ctx context.Context, todo model.Todo) error {
	return s.repo.Add(todo)
}
