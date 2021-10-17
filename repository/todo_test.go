package repository

import (
	"app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoRepository_GetTodos(t *testing.T) {
	t.Run("Initially length has to be 0", func(t *testing.T) {
		todoRepository := &todoRepository{
			todos: make([]model.Todo, 0),
		}
		todos, error := todoRepository.GetAll()
		assert.Nil(t, error)
		assert.Equal(t, 0, len(todos))
	})

	t.Run("After adding one todo length has to be 1", func(t *testing.T) {
		todoRepository := &todoRepository{
			todos: make([]model.Todo, 0),
		}
		todoRepository.todos = append(todoRepository.todos, model.Todo{
			Id:   1,
			Text: "Test",
		})
		todos, error := todoRepository.GetAll()
		assert.Nil(t, error)
		assert.Equal(t, 1, len(todos))
		assert.Equal(t, 1, todos[0].Id)
		assert.Equal(t, "Test", todos[0].Text)
	})
}

func TestTodoRepository_Add(t *testing.T) {
	todoRepository := &todoRepository{
		todos: make([]model.Todo, 0),
	}
	todoRepository.Add(model.Todo{
		Id:   1,
		Text: "Test",
	})
	assert.Equal(t, 1, len(todoRepository.todos))
	assert.Equal(t, 1, todoRepository.todos[0].Id)
	assert.Equal(t, "Test", todoRepository.todos[0].Text)
}
