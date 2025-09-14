-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE is_active = true
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
    email, username, password_hash, full_name
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE($2, email),
    username = COALESCE($3, username),
    full_name = COALESCE($4, full_name),
    is_active = COALESCE($5, is_active)
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;