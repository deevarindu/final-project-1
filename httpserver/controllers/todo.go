package controllers

import (
	"net/http"

	"github.com/deevarindu/final-project-1/httpserver/controllers/params"
	"github.com/deevarindu/final-project-1/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoController struct {
	svc *services.TodoSvc
}

func NewTodoController(svc *services.TodoSvc) *TodoController {
	return &TodoController{
		svc: svc,
	}
}

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.Todo
// @Router /todos [get]
func (t *TodoController) GetTodos(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := t.svc.GetTodos()
	WriteJsonResponse(ctx, response)
}

// CreateTodo godoc
// @Summary Create new todo
// @Description Create new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body params.TodoCreateRequest true "Create Todo"
// @Success 201 {array} models.Todo
// @Router /todos [post]
func (t *TodoController) CreateTodo(ctx *gin.Context) {
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

	response := t.svc.CreateTodo(&todo)
	WriteJsonResponse(ctx, response)
}

// GetTodoById godoc
// @Summary Get todo by id
// @Description Get todo by id
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos/{id} [get]
func (t *TodoController) GetTodoById(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	id := ctx.Param("id")
	response := t.svc.GetTodoById(id)
	WriteJsonResponse(ctx, response)
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
func (t *TodoController) UpdateTodo(ctx *gin.Context) {
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

	response := t.svc.UpdateTodo(&todo)
	WriteJsonResponse(ctx, response)
}

// DeleteTodo godoc
// @Summary Delete todo
// @Description Delete todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos/{id} [delete]
func (t *TodoController) DeleteTodo(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	id := ctx.Param("id")
	response := t.svc.DeleteTodo(id)
	WriteJsonResponse(ctx, response)
}
