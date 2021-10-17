package service

import (
	"app/model"
	"app/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoService_GetTodos(t *testing.T) {
	repo := repository.NewTodoRepository()
	service := NewTodoService(repo)
	todos, err := service.GetTodos(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))

	repo.Add(model.Todo{
		Id:   1,
		Text: "test",
	})

	todos, err = service.GetTodos(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(todos))
}

func TestTodoRepository_Add(t *testing.T) {
	repo := repository.NewTodoRepository()
	service := NewTodoService(repo)

	todo := model.Todo{
		Id:   1,
		Text: "test",
	}

	err := service.AddTodo(context.TODO(), todo)
	assert.Nil(t, err)

	todos, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(todos))
}
