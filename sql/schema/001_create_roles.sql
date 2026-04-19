-- +goose Up

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
)

-- +goose Down

DROP TABLE roles;