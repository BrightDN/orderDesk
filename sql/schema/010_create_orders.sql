-- +goose Up

CREATE TABLE orders (
    id TEXT PRIMARY KEY,

    supplier_id INTEGER NOT NULL,
    employee_id INTEGER NOT NULL,
    company_id INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    FOREIGN KEY (supplier_id, company_id)
        REFERENCES suppliers(id, company_id)
        ON DELETE CASCADE,

    FOREIGN KEY (employee_id, company_id)
        REFERENCES employees(id, company_id)
        ON DELETE RESTRICT
);

-- +goose Down

DROP TABLE IF EXISTS orders;
