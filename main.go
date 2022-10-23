package main

import (
	"log"
	"net/http"

	"github.com/deevarindu/final-project-1/httpserver/controllers"
	"github.com/gorilla/mux"
)

// @title Final Project 1 Todos API Documentation
// @version 1.0
// @description This is the simple documentation for the Final Project 1 Todos API.
// @termsOfService http://swagger.io/terms/
// @contact.email deevarindu@gmail.com
// @contact.name Deeva Rindu
// @host localhost:5000
// @BasePath /

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ping", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controllers.GetTodoById).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	// router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
