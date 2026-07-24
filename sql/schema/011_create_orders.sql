-- +goose Up

CREATE TABLE orders (
    id TEXT PRIMARY KEY,

    supplier_id INTEGER NOT NULL,
    employee_id INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    FOREIGN KEY (supplier_id)
        REFERENCES suppliers(id)
        ON DELETE CASCADE,
    
    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE CASCADE,

    FOREIGN KEY (employee_id)
        REFERENCES employees(id)
        ON DELETE RESTRICT
);

-- +goose Down

DROP TABLE IF EXISTS orders;