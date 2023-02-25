-- name: CreateTodo :one
INSERT INTO todos ( name , description , done )
    VALUES ($1, $2, $3) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
    WHERE id = $1 LIMIT 1;

-- name: GetAllTodo :many
SELECT * FROM todos;

-- name: UpdateTodo :one
UPDATE todos
SET name = $2 , description = $3 , done = #4
    WHERE id = $1  RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
    WHERE id = $1;