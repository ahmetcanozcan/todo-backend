package repository

import (
	"app/model"
	"sync"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	Add(todo model.Todo) error
}

type todoRepository struct {
	todos []model.Todo
	mu    sync.Mutex
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{
		todos: make([]model.Todo, 0),
		mu:    sync.Mutex{},
	}
}

func (t *todoRepository) GetAll() ([]model.Todo, error) {
	return t.todos, nil
}

func (t *todoRepository) Add(todo model.Todo) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.todos = append(t.todos, todo)
	return nil
}
