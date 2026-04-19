-- +goose Up

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    company_id INTEGER NOT NULL,
    
    quantity INTEGER NOT NULL,
    name_at_order TEXT NOT NULL,
    price_at_order NUMERIC(10, 2),

    FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT
);

-- +goose Down

DROP IF EXISTS order_items;