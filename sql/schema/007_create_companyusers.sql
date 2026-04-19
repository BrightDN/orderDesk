-- +goose Up

CREATE TABLE company_users (
    company_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,

    display_name TEXT NOT NULL,
    joined_at TIMESTAMPTZ DEFAULT now(),

    PRIMARY KEY (company_id, user_id) ,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE RESTRICT,
);

-- +goose Down

DROP TABLE IF EXISTS company_users;