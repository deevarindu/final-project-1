package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/deevarindu/final-project-1/httpserver/controllers/params"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} views.Response
// @Router /todos [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "ok",
	})
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
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo params.TodoCreateRequest
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": todo,
	})
}
