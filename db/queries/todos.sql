-- name: CreateTodo :exec
INSERT INTO todos (title, description, priority, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id;  -- Return the generated ID

-- name: GetTodoByID :one
SELECT * FROM todos WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos;

-- name: UpdateTodo :exec
UPDATE todos SET title = $2, description = $3, priority = $4, updated_at = NOW() WHERE id = $1;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;
