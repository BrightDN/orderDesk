-- +goose Up

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,

    name TEXT NOT NULL
);

-- +goose Down

DROP TABLE IF EXISTS users;