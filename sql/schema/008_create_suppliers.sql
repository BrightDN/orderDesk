-- +goose Up

CREATE TABLE suppliers (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    contact TEXT,

    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    company_id INTEGER NOT NULL,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT,

    CONSTRAINT suppliers_company_name_unique
        UNIQUE (company_id, name),

    CONSTRAINT suppliers_company_email_unique
        UNIQUE (company_id, email)
);

-- +goose Down

DROP TABLE IF EXISTS suppliers;