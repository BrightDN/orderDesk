-- +goose Up

CREATE TABLE companies (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  deleted_at TIMESTAMPTZ DEFAULT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT companies_id_unique UNIQUE (id)
);

CREATE UNIQUE INDEX companies_active_name_unique
  ON companies (name)
  WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX companies_active_email_unique
  ON companies (email)
  WHERE deleted_at IS NULL;

-- +goose StatementBegin
CREATE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER companies_set_updated_at
BEFORE UPDATE ON companies
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- +goose Down

DROP TABLE IF EXISTS companies;
DROP FUNCTION IF EXISTS set_updated_at();
