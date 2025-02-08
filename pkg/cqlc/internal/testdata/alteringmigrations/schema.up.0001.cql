CREATE TABLE users (
    user_id    UUID PRIMARY KEY,
    username   TEXT,
    email      TEXT,
    created_at TIMESTAMP
);

CREATE TABLE logins (
    used_id    UUID PRIMARY KEY,
    login_time TIMESTAMP
);
