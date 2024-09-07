-- name: CreateUser :exec
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id;  -- Return the generated ID

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users SET email = $2, password = $3, updated_at = NOW() WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
