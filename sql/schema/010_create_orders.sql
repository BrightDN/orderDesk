-- +goose Up

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    supplier_id INTEGER NOT NULL,
    company_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    FOREIGN KEY (supplier_id)
        REFERENCES suppliers(id)
        ON DELETE RESTRICT,
    
    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE RESTRICT
);

-- +goose Down

DROP TABLE IF EXISTS orders;