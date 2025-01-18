-- name: CreateUser :exec
INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);

-- name: FindUser :one
SELECT * FROM users WHERE user_id = ? LIMIT 1;

-- name: ListUserIDs :many
SELECT user_id FROM users;

-- name: ListUsers :many
SELECT * FROM users;

-- name: ListUserPosts :many
SELECT * FROM posts WHERE user_id = ? AND created_at = ?;