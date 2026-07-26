package repositories

import (
	"backend/internal/app/database"
	"backend/internal/app/entities"
	"database/sql"
)

type Todo struct {
	database.BaseSQLRepository[entities.Todo]
}

func NewTodoRepository(db *sql.DB) *Todo {
	return &Todo{
		BaseSQLRepository: database.BaseSQLRepository[entities.Todo]{DB: db},
	}
}

func mapTodo(rows *sql.Row, t *entities.Todo) error {
	return rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
}

func mapTodos(rows *sql.Rows, t *entities.Todo) error {
	return rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
}

func (r *Todo) FindByID(id int, userID int) (*entities.Todo, error) {
	return r.SelectSingle(
		mapTodo,
		"SELECT id, title, description, completed, user_id, created_at, updated_at FROM todos WHERE id = $1 AND user_id = $2",
		id,
		userID,
	)
}

func (r *Todo) GetAllTodos(userID int) ([]*entities.Todo, error) {
	return r.SelectMultiple(
		mapTodos,
		"SELECT id, title, description, completed, user_id, created_at, updated_at FROM todos WHERE user_id = $1",
		userID,
	)
}

func (r *Todo) GetCompletedTodos(userID int) ([]*entities.Todo, error) {
	return r.SelectMultiple(
		mapTodos,
		"SELECT id, title, description, completed, user_id, created_at, updated_at FROM todos WHERE user_id = $1 AND completed = true",
		userID,
	)
}

func (r *Todo) GetIncompleteTodos(userID int) ([]*entities.Todo, error) {
	return r.SelectMultiple(
		mapTodos,
		"SELECT id, title, description, completed, user_id, created_at, updated_at FROM todos WHERE user_id = $1 AND completed = false",
		userID,
	)
}

func (r *Todo) CreateTodo(todo *entities.Todo) error {
	id, err := r.Insert(
		"INSERT INTO todos (title, description, completed, user_id) VALUES ($1, $2, $3, $4)",
		todo.Title,
		todo.Description,
		todo.Completed,
		todo.UserID,
	)
	todo.ID = id
	return err
}

func (r *Todo) UpdateTodo(todo *entities.Todo) error {
	_, err := r.ExecuteQuery(
		"UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4 AND user_id = $5",
		todo.Title,
		todo.Description,
		todo.Completed,
		todo.ID,
		todo.UserID,
	)
	return err
}

func (r *Todo) DeleteTodo(id int, userID int) error {
	_, err := r.ExecuteQuery(
		"DELETE FROM todos WHERE id = $1 AND user_id = $2",
		id,
		userID,
	)
	return err
}
