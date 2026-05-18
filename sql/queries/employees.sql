-- name: GetCompanyEmployees :many
SELECT *
FROM company_users
WHERE company_id = $1;