-- +goose Up

CREATE TABLE invites (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,

    company_id INTEGER NOT NULL,
    invite_type TEXT NOT NULL,

    token TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,
    used_at TIMESTAMPTZ DEFAULT NULL,

    role_id INTEGER NOT NULL default 2,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE CASCADE,

    FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE RESTRICT,

    CONSTRAINT unique_invite_token
        UNIQUE (token),

    CONSTRAINT valid_expiration
        CHECK (expires_at > created_at),

    CONSTRAINT valid_invite_type
        CHECK (invite_type IN ('company', 'employee'))
);

CREATE TRIGGER invites_set_updated_at
BEFORE UPDATE ON invites
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- +goose Down

DROP TABLE IF EXISTS invites;
