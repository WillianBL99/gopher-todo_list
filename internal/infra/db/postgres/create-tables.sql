CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    list VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP,
    deleted_at TIMESTAMP
);





