-- +goose Up

CREATE TABLE order_mails (
    id SERIAL PRIMARY KEY,

    supplier_id INT NOT NULL,
    subject TEXT NOT NULL,
    mail_content TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL default now(),
    updated_at TIMESTAMPTZ NOT NULL default now(),

    FOREIGN KEY (supplier_id)
        REFERENCES suppliers(id)
        ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS order_mails;