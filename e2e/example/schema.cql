CREATE TABLE users
(
    user_id    UUID PRIMARY KEY,
    username   TEXT,
    email      TEXT,
    created_at TIMESTAMP
);