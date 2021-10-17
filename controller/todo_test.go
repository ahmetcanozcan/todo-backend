package controller

import (
	"app/model"
	"app/repository"
	"app/service"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoController_GetTodos(t *testing.T) {
	req := httptest.NewRequest("GET", "/todos", nil)

	r := repository.NewTodoRepository()
	s := service.NewTodoService(r)
	c := NewTodoController(s)

	t.Run("Returns empty list", func(t *testing.T) {
		w := httptest.NewRecorder()
		c.GetTodos(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "[]\n", w.Body.String())
	})

	t.Run("Returns list of todos", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.Add(model.Todo{
			Text: "test",
		})
		c.GetTodos(w, req)
		assert.Equal(t, 200, w.Code)

		var todos []model.Todo
		err := json.Unmarshal(w.Body.Bytes(), &todos)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(todos))
		assert.Equal(t, "test", todos[0].Text)
	})

}
