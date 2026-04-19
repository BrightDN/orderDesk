-- +goose Up

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    permission TEXT NOT NULL
);

-- +goose Down

DROP TABLE permissions;