-- +goose Up

CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    email TEXT NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- +goose Down

DROP TABLE IF EXISTS companies;