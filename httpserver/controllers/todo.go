package controllers

import (
	"net/http"

	"github.com/deevarindu/final-project-1/httpserver/controllers/params"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.Todo
// @Router /todos [get]
func GetTodos(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

}

// CreateTodo godoc
// @Summary Create new todo
// @Description Create new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body params.TodoCreateRequest true "Create Todo"
// @Success 200 {array} models.Todo
// @Router /todos [post]
func CreateTodo(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var todo params.TodoCreateRequest
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

// GetTodoById godoc
// @Summary Get todo by id
// @Description Get todo by id
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos/{id} [get]
func GetTodoById(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

}

// UpdateTodo godoc
// @Summary Update todo
// @Description Update todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body params.TodoUpdateRequest true "Update Todo"
// @Success 200 {array} models.Todo
// @Router /todos/{id} [put]
func UpdateTodo(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var todo params.TodoUpdateRequest
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

// DeleteTodo godoc
// @Summary Delete todo
// @Description Delete todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos/{id} [delete]
func DeleteTodo(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
}
