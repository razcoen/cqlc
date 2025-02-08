-- name: CreateUser :exec
INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);

-- name: FindUser :one
SELECT * FROM users WHERE user_id = ? LIMIT 1;

-- name: RecordLogin :exec
INSERT INTO logins (user_id, login_time) VALUES (?, ?);

-- name: ListLogins :many
SELECT login_time FROM logins WHERE user_id = ?;
