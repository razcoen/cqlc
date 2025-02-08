ALTER TABLE users ADD last_login DATE;
ALTER TABLE users DROP created_at;
DROP TABLE logins;
