-- name: CreateUser :exec
INSERT INTO users (email, password)
VALUES (?, ?)
RETURNING *;  -- Return the generated ID

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT id, email, password FROM users WHERE email = ?;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
SET email = ?, password = ?, updated_at = NOW()
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
