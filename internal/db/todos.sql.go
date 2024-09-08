// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: todos.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTodo = `-- name: CreateTodo :exec
INSERT INTO todos (title, description, priority, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateTodoParams struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Priority    string      `json:"priority"`
	UserID      pgtype.Int8 `json:"user_id"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) error {
	_, err := q.db.Exec(ctx, createTodo,
		arg.Title,
		arg.Description,
		arg.Priority,
		arg.UserID,
	)
	return err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1 AND user_id = $2
`

type DeleteTodoParams struct {
	ID     int64       `json:"id"`
	UserID pgtype.Int8 `json:"user_id"`
}

func (q *Queries) DeleteTodo(ctx context.Context, arg DeleteTodoParams) error {
	_, err := q.db.Exec(ctx, deleteTodo, arg.ID, arg.UserID)
	return err
}

const getTodoByID = `-- name: GetTodoByID :one

SELECT id, title, description, priority, created_at, updated_at, user_id FROM todos WHERE id = $1 AND user_id = $2
`

type GetTodoByIDParams struct {
	ID     int64       `json:"id"`
	UserID pgtype.Int8 `json:"user_id"`
}

// Return the generated ID
func (q *Queries) GetTodoByID(ctx context.Context, arg GetTodoByIDParams) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodoByID, arg.ID, arg.UserID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Priority,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, description, priority, created_at, updated_at, user_id FROM todos WHERE user_id = $1
`

func (q *Queries) ListTodos(ctx context.Context, userID pgtype.Int8) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodos, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Priority,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :exec
UPDATE todos SET title = $2, description = $3, priority = $4, updated_at = NOW() WHERE id = $1 AND user_id = $5
`

type UpdateTodoParams struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Priority    string      `json:"priority"`
	UserID      pgtype.Int8 `json:"user_id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.Exec(ctx, updateTodo,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Priority,
		arg.UserID,
	)
	return err
}
