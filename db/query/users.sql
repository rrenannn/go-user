-- name: CreateUser :one
INSERT INTO users(name, email, password, status, created_at)
VALUES ($1, $2, $3, $4, now())
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 AND status = TRUE;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 AND status = TRUE;

-- name: ResetPassword :exec
UPDATE users
SET password = $3
WHERE email = $1 AND status = $2;