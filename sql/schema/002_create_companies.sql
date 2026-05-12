-- +goose Up

CREATE TABLE companies (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  deleted_at TIMESTAMPTZ DEFAULT NULL,
  CONSTRAINT unique_name UNIQUE (name),
  CONSTRAINT unique_email UNIQUE (email)
);

-- +goose Down

DROP TABLE IF EXISTS companies;