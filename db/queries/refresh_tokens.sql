-- name: StoreRefreshToken :exec
INSERT INTO refresh_tokens (token, user_id, expires_at, ip_address, device)
VALUES ($1, $2, $3, $4, $5)
RETURNING id; -- Return the generated ID

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE id = $1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens WHERE id = $1;

-- name: DeleteRefreshTokensByUserId :exec
DELETE FROM refresh_tokens WHERE user_id = $1;