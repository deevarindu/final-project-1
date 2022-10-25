package services

import (
	"fmt"
	"strings"

	"github.com/deevarindu/final-project-1/httpserver/controllers/params"
	"github.com/deevarindu/final-project-1/httpserver/repositories"
	"github.com/deevarindu/final-project-1/httpserver/repositories/models"
	"github.com/deevarindu/final-project-1/httpserver/views"
)

type TodoSvc struct {
	repo repositories.TodoRepository
}

func NewTodoSvc(repo repositories.TodoRepository) *TodoSvc {
	return &TodoSvc{
		repo: repo,
	}
}

func (t *TodoSvc) GetTodos() *views.Response {
	todos, err := t.repo.GetTodos()
	if err != nil {
		return views.BadRequestResponse(err)
	}
	return views.SuccessResponse(todos, "Get all todos successfully")
}

func (t *TodoSvc) GetTodoById(id string) *views.Response {
	todo, err := t.repo.GetTodoById(id)
	if err != nil {
		return views.BadRequestResponse(err)
	}
	return views.SuccessResponse(todo, "Get todo successfully")
}

func (t *TodoSvc) CreateTodo(req *params.TodoCreateRequest) *views.Response {
	todo := ParseTodoCreateRequest(req)
	todos, _ := t.repo.GetTodos()
	id := "1"
	if todos != nil {
		if len(*todos) > 0 {
			id = fmt.Sprintf("%d", len(*todos)+1)
		}
	}

	todo.ID = &id

	err := t.repo.CreateTodo(todo)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return views.DataConflictResponse(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessResponse(todo, "Create todo successfully")
}

func ParseTodoCreateRequest(todo *params.TodoCreateRequest) *models.Todo {
	return &models.Todo{
		Title:  todo.Title,
		Status: todo.Status,
	}
}

func (t *TodoSvc) UpdateTodo(req *params.TodoUpdateRequest) *views.Response {
	todo := ParseTodoUpdateRequest(req)
	err := t.repo.UpdateTodo(todo)
	if err != nil {
		return views.InternalServerError(err)
	}
	return views.SuccessResponse(todo, "Update todo successfully")
}

func ParseTodoUpdateRequest(todo *params.TodoUpdateRequest) *models.Todo {
	return &models.Todo{
		Title:  todo.Title,
		Status: todo.Status,
	}
}

func (t *TodoSvc) DeleteTodo(id string) *views.Response {
	err := t.repo.DeleteTodo(id)
	if err != nil {
		return views.InternalServerError(err)
	}
	return views.SuccessDeleteResponse("Delete todo successfully")
}
