package main

import (
	"app/controller"
	"app/repository"
	"app/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	repo := repository.NewTodoRepository()
	service := service.NewTodoService(repo)
	controller := controller.NewTodoController(service)

	router := mux.NewRouter()

	router.Methods("GET").Path("/todos").HandlerFunc(controller.GetTodos)
	router.Methods("POST").Path("/todos").HandlerFunc(controller.AddTodo)

	server := http.Server{
		Handler: router,
		Addr:    ":8000",
	}
	server.ListenAndServe()
}
