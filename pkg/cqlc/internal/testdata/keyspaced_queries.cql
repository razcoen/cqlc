-- name: CreateUser :exec
INSERT INTO auth.users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);

-- name: CreateUsers :batch
INSERT INTO auth.users (user_id, username, email, created_at) VALUES (?, ?, ?, ?);

-- name: DeleteUser :exec
DELETE FROM auth.users WHERE user_id = ?;

-- name: DeleteUsers :batch
DELETE FROM auth.users WHERE user_id = ?;

-- name: FindUser :one
SELECT * FROM auth.users WHERE user_id = ? LIMIT 1;

-- name: FindUsers :many
SELECT * FROM auth.users WHERE email = ? ALLOW FILTERING;
