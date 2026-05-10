-- +goose Up

CREATE TABLE invites (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,

    company_id INTEGER NOT NULL,
    invite_type TEXT NOT NULL,

    token TEXT NOT NULL,

    created_at TIMESTAMPTZ DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,
    used_at TIMESTAMPTZ DEFAULT NULL,

    UNIQUE (email, company_id),

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE RESTRICT,

    CONSTRAINT unique_invite_token
        UNIQUE (token),

    CONSTRAINT valid_invite_type
        CHECK (invite_type IN ('company', 'employee'))
);

-- +goose Down

DROP TABLE IF EXISTS invites;