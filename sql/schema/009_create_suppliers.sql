-- +goose Up

CREATE TABLE suppliers (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    contact TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    is_active BOOLEAN NOT NULL DEFAULT 'true',

    company_id INTEGER NOT NULL,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE CASCADE,

    CONSTRAINT suppliers_company_name_unique
        UNIQUE (company_id, name),

    CONSTRAINT suppliers_company_email_unique
        UNIQUE (company_id, email)
);

-- +goose Down

DROP TABLE IF EXISTS suppliers;