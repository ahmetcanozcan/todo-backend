package controller

import (
	"app/model"
	"app/service"
	"encoding/json"
	"net/http"
)

type todoController struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) *todoController {
	return &todoController{
		service: service,
	}
}

func (t *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todos, err := t.service.GetTodos(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (t *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = t.service.AddTodo(ctx, todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
