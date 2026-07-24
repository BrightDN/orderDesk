-- +goose Up

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    company_id INTEGER NOT NULL,
    user_id INTEGER,
    role_id INTEGER NOT NULL,

    display_name TEXT NOT NULL,
    
    joined_at TIMESTAMPTZ DEFAULT now(),
    left_at TIMESTAMPTZ DEFAULT NULL,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE CASCADE,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE SET NULL,

    FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE RESTRICT,

    CONSTRAINT employees_company_user_unique
        UNIQUE (company_id, user_id)
);

-- +goose Down

DROP TABLE IF EXISTS employees;