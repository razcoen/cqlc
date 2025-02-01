-- name: CreateUser :exec
INSERT INTO users (user_id, username, email, created_at) VALUES (?, ?, ?);

-- name: CreateUsers :batch
INSERT INTO users (user_id, username, email, created_at VALUES (?, ?, ?, ?);
