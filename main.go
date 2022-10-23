package main

import (
	"fmt"

	"github.com/deevarindu/final-project-1/config"
	"github.com/deevarindu/final-project-1/routes"
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

	if db != nil {
		fmt.Println("Connected to database")
	}

	app := routes.RegisterRoutes()
	app.Run(":5000")
}
