-- +goose Up

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    supplier_id INTEGER NOT NULL,
    company_id INTEGER NOT NULL,
    price NUMERIC(10, 2),
    deleted_at TIMESTAMPTZ,

    FOREIGN KEY (supplier_id)
        REFERENCES suppliers(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT
)

-- +goose Down

DROP TABLE IF EXISTS products;