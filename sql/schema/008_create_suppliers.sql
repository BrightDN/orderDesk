-- +goose Up

CREATE TABLE suppliers (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    contact TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    is_active BOOLEAN NOT NULL DEFAULT 'true',

    company_id INTEGER NOT NULL,

    FOREIGN KEY (company_id)
        REFERENCES companies(id)
        ON DELETE CASCADE,

    CONSTRAINT suppliers_id_company_unique
        UNIQUE (id, company_id)
);

CREATE UNIQUE INDEX suppliers_active_company_name_unique
  ON suppliers (company_id, name)
  WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX suppliers_active_company_email_unique
  ON suppliers (company_id, email)
  WHERE deleted_at IS NULL;

CREATE TRIGGER suppliers_set_updated_at
BEFORE UPDATE ON suppliers
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- +goose Down

DROP TABLE IF EXISTS suppliers;
