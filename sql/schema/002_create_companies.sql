-- +goose Up

CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
)

-- +goose Down

DROP TABLE companies;