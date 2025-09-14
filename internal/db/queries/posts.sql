-- name: GetPost :one
SELECT p.*, u.username, u.email
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT p.*, u.username, u.email
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.status = 'published'
ORDER BY p.published_at DESC
LIMIT $1 OFFSET $2;

-- name: ListUserPosts :many
SELECT * FROM posts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreatePost :one
INSERT INTO posts (
    user_id, title, content, status
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET
    title = COALESCE($2, title),
    content = COALESCE($3, content),
    status = COALESCE($4, status),
    published_at = CASE
        WHEN $4 = 'published' AND status != 'published' THEN CURRENT_TIMESTAMP
        ELSE published_at
    END
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

-- name: CountPosts :one
SELECT COUNT(*) FROM posts
WHERE status = $1;