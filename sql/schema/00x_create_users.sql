-- +goose Up

CREATE TABLE users(
    id UUID PRIMARY KEY NOT NULL,
    created_at DATETIME NOT NULL DEFAULT now(),
    updated_at DATETIME NOT NULL DEFAULT now(),

    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    name TEXT NOT NULL,

    company_id UUID NOT NULL,
    role_id integer NOT NULL,
    permissions TEXT[] NOT NULL DEFAULT '{}'

    FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE users;