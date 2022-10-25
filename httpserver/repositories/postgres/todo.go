package postgres

import (
	"database/sql"

	"github.com/deevarindu/final-project-1/httpserver/repositories"
	"github.com/deevarindu/final-project-1/httpserver/repositories/models"
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) repositories.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (t *todoRepository) GetTodos() (*[]models.Todo, error) {
	query := `SELECT id, title, status FROM todos`

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo := models.Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return &todos, nil
}

func (t *todoRepository) CreateTodo(todo *models.Todo) error {
	query := `
		INSERT INTO todos (id, title, status) VALUES
		($1, $2, $3)
	`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.ID, todo.Title, todo.Status)

	return err
}

func (t *todoRepository) GetTodoById(id string) (*models.Todo, error) {
	query := `SELECT id, title, status FROM todos WHERE id = $1`

	row := t.db.QueryRow(query, id)

	todo := models.Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *todoRepository) UpdateTodo(todo *models.Todo) error {
	query := `
		UPDATE todos SET title = $1, status = $2 WHERE id = $3
	`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Status, todo.ID)

	return err
}

func (t *todoRepository) DeleteTodo(id string) error {
	query := `DELETE FROM todos WHERE id = $1`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
