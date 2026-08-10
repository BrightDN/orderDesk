-- +goose Up

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id TEXT NOT NULL,
    product_id INTEGER,
    
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    name_at_order TEXT NOT NULL,

    FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE SET NULL
);

-- +goose Down

DROP TABLE IF EXISTS order_items;
