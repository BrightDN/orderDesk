-- name: CreateCompany :one

INSERT INTO companies (name, email)
VALUES (
    $1,
    $2
)RETURNING *;

-- name: GetCompanies :many
SELECT * FROM companies;