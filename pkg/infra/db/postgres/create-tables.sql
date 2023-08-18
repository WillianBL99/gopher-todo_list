CREATE DATABASE todo_list;
\c todo_list;
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP,
    deleted_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    description VARCHAR(255) NOT NULL
);





