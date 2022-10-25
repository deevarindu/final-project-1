package repositories

import (
	"github.com/deevarindu/final-project-1/httpserver/repositories/models"
)

type TodoRepository interface {
	GetTodos() (*[]models.Todo, error)
	CreateTodo(todo *models.Todo) error
	GetTodoById(id string) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id string) error
}
