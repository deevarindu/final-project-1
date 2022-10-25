package routes

import (
	"github.com/deevarindu/final-project-1/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	todo   *controllers.TodoController
}

func NewRouter(router *gin.Engine, todo *controllers.TodoController) *Router {
	return &Router{
		router: router,
		todo:   todo,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/todos", r.todo.GetTodos)
	r.router.POST("/todos", r.todo.CreateTodo)
	r.router.GET("/todos/:id", r.todo.GetTodoById)
	r.router.PUT("/todos/:id", r.todo.UpdateTodo)
	r.router.DELETE("/todos/:id", r.todo.DeleteTodo)
	r.router.Run(port)
}
