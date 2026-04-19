-- +goose Up

CREATE TABLE invites (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,

    company_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,

    token TEXT NOT NULL UNIQUE,
    invited_by INTEGER NOT NULL,

    created_at TIMESTAMPTZ DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,

    UNIQUE (email, company_id),

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE RESTRICT,

    FOREIGN KEY (invited_by)
        REFERENCES company_users(id)
        ON DELETE SET NULL
);

-- +goose Down

DROP IF EXISTS invites;