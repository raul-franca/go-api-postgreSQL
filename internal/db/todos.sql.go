// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: todos.sql

package db

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos ( name , description , done )
    VALUES ($1, $2, $3) RETURNING id, name, description, done, created_at
`

type CreateTodoParams struct {
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	Done        sql.NullBool   `json:"done"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.Name, arg.Description, arg.Done)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Done,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
    WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getAllTodo = `-- name: GetAllTodo :many
SELECT id, name, description, done, created_at FROM todos
`

func (q *Queries) GetAllTodo(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getAllTodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Todo{}
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Done,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodo = `-- name: GetTodo :one
SELECT id, name, description, done, created_at FROM todos
    WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Done,
		&i.CreatedAt,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET name = $2 , description = $3 , done = #4
    WHERE id = $1  RETURNING id, name, description, done, created_at
`

type UpdateTodoParams struct {
	ID          int32          `json:"id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.ID, arg.Name, arg.Description)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Done,
		&i.CreatedAt,
	)
	return i, err
}
