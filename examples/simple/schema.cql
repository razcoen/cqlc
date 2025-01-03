-- Keyspace creation
CREATE KEYSPACE test_keyspace
            WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};

-- Switch to keyspace
USE test_keyspace;

-- Create a table for users
CREATE TABLE users
(
    user_id    UUID PRIMARY KEY,
    username   TEXT,
    email      TEXT,
    created_at TIMESTAMP
) WITH default_time_to_live = 0;

-- Create a table for posts with clustering columns
CREATE TABLE posts
(
    post_id    UUID,
    user_id    UUID,
    content    TEXT,
    created_at TIMESTAMP,
    PRIMARY KEY (user_id, created_at)
) WITH CLUSTERING ORDER BY (created_at DESC);

-- Create a table with collections
CREATE TABLE user_profiles
(
    user_id    UUID PRIMARY KEY,
    first_name TEXT,
    last_name  TEXT,
    interests  LIST<TEXT>,
    settings   MAP<TEXT, TEXT>
);

-- Create a table with a composite key
CREATE TABLE orders
(
    order_id   UUID,
    user_id    UUID,
    product_id UUID,
    quantity   INT,
    order_date TIMESTAMP,
    PRIMARY KEY ((order_id, user_id), product_id)
) WITH CLUSTERING ORDER BY (product_id ASC);

-- Create a user-defined type
CREATE TYPE address
    (
        street   TEXT,
        city     TEXT,
        state    TEXT,
        zip_code INT
    );

-- Create a table that uses a UDT
CREATE TABLE customer_addresses
(
    customer_id  UUID PRIMARY KEY,
    address_list LIST<FROZEN<address>>
);

-- Create a table with secondary indexes
CREATE TABLE products
(
    product_id UUID PRIMARY KEY,
    name       TEXT,
    category   TEXT,
    price      DECIMAL
);

CREATE INDEX ON products (category);

-- Create a table with a TTL and default options
CREATE TABLE temporary_sessions
(
    session_id UUID PRIMARY KEY,
    user_id    UUID,
    created_at TIMESTAMP
) WITH default_time_to_live = 3600; -- 1 hour
