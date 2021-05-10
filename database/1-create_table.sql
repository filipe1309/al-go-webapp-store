-- CREATE DATABASE go_store ENCODING 'UTF8';
-- DROP DATABASE IF EXISTS go_store;
CREATE DATABASE go_store;
\c go_store;
-- CREATE schema go_store;
-- SET search_path = go_store;
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR,
    price DECIMAL,
    quantity INTEGER
)