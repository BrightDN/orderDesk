-- +goose Up

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    supplier_id INTEGER NOT NULL,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    FOREIGN KEY (supplier_id)   
        REFERENCES suppliers(id)
        ON DELETE RESTRICT,

    CONSTRAINT products_supplier_name_unique
        UNIQUE (supplier_id, name)
);

-- +goose Down

DROP TABLE IF EXISTS products;