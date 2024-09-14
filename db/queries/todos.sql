-- name: CreateTodo :exec
INSERT INTO todos (title, description, priority, user_id)
VALUES (?, ?, ?, ?)
RETURNING id;  -- Return the generated ID

-- name: GetTodoByID :one
SELECT * FROM todos WHERE id = ? AND user_id = ?;

-- name: ListTodos :many
SELECT * FROM todos WHERE user_id = ?;

-- name: UpdateTodo :exec
UPDATE todos
SET title = ?, description = ?, priority = ?, complete = ?, updated_at = NOW()
WHERE id = ? AND user_id = ?
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ? AND user_id = ?;
