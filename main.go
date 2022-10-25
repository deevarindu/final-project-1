package main

import (
	"github.com/deevarindu/final-project-1/config"
	"github.com/deevarindu/final-project-1/httpserver/controllers"
	"github.com/deevarindu/final-project-1/httpserver/repositories/postgres"
	"github.com/deevarindu/final-project-1/httpserver/services"
	"github.com/deevarindu/final-project-1/routes"
	"github.com/gin-gonic/gin"
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
	db, err := config.CreateConnection()
	if err != nil {
		panic(err)
	}

	todoRepository := postgres.NewTodoRepository(db)
	todoSvc := services.NewTodoSvc(todoRepository)
	todoHandler := controllers.NewTodoController(todoSvc)

	router := gin.Default()

	app := routes.NewRouter(router, todoHandler)
	app.Start(":5000")
}
