-- +goose Up

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    permission TEXT NOT NULL UNIQUE
);

-- +goose Down

DROP TABLE IF EXISTS permissions;