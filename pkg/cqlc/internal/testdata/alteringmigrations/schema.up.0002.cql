ALTER TABLE users ADD last_login TIME;
ALTER TABLE users DROP created_at;
DROP TABLE logins;
