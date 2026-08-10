-- +goose Up

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  name TEXT NOT NULL,

  is_admin BOOLEAN NOT NULL DEFAULT 'false',

  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ DEFAULT NULL
);

-- +goose Down

DROP TABLE IF EXISTS users;