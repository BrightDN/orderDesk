-- name: CreateCompany :one
INSERT INTO
  companies (name, email)
VALUES
  ($1, $2)
RETURNING
  *;

-- name: GetCompanies :many
SELECT
  *
FROM
  companies;

-- name: GetCompany :one
SELECT
  *
FROM
  companies
WHERE
  id = $1;

-- name: DeleteCompany :exec
UPDATE companies
SET
  deleted_at = NOW(),
  updated_at = NOW()
WHERE
  id = $1
  AND deleted_at IS NULL;

-- name: UpdateCompany :exec
UPDATE companies
  SET
    name = $2,
    email = $3,
    updated_at = NOW()
WHERE id = $1;