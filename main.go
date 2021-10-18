package main

import (
	"app/controller"
	"app/repository"
	"app/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	repo := repository.NewTodoRepository()
	service := service.NewTodoService(repo)
	controller := controller.NewTodoController(service)

	router := mux.NewRouter()

	router.Methods("GET").Path("/todos").HandlerFunc(controller.GetTodos)
	router.Methods("POST").Path("/todos").HandlerFunc(controller.AddTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Server is listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(headersOk, originsOk, methodsOk)(router)))

}
