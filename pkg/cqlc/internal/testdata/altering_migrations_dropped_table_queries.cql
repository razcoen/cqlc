-- name: CreateUser :exec
INSERT INTO users (user_id, username, email, last_login) VALUES (?, ?, ?, ?);

-- name: FindUser :one
SELECT * FROM users WHERE user_id = ? LIMIT 1;

-- name: ListLogins :many
SELECT login_time FROM logins WHERE user_id = ?;
