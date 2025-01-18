-- name: CreateUser :exec
INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);

-- name: FindUser :one
SELECT * FROM users WHERE user_id = ? LIMIT 1;

-- name: FindUsernameByUserID :one
SELECT username FROM users WHERE user_id = ? LIMIT 1;

-- name: FindUsers :many
SELECT * FROM users WHERE email = ?;
