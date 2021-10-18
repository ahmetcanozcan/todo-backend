package main

import (
	"app/controller"
	"app/repository"
	"app/service"
	"log"
	"net/http"
	"os"
	"time"

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
		port = "8901"
	}
	ch := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)
	addr := ":" + port

	s := http.Server{
		Addr:         addr,
		Handler:      ch(router),
		ErrorLog:     log.Default(),
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	log.Fatal(s.ListenAndServe())
}
