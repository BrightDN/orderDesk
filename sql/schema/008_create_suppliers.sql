-- +goose Up

CREATE TABLE suppliers (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,

    deleted_at TIMESTAMPTZ,

    company_id INTEGER NOT NULL,


    UNIQUE (company_id, name),
    UNIQUE (company_id, email),

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT
);

-- +goose Down

DROP TABLE IF EXISTS suppliers;