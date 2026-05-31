-- name: GetCompanyEmployees :many
SELECT
  company_users.display_name,
  company_users.user_id,
  company_users.company_id,
  company_users.id AS employee_id,
  users.email AS email,
  roles.name AS role
FROM
  company_users
  INNER JOIN users ON company_users.user_id = users.id
  INNER JOIN roles ON company_users.role_id = roles.id
WHERE
  company_id = $1;

-- name: CreateCompanyEmployee :one
INSERT INTO company_users (company_id, user_id, role_id, display_name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEmployee :one
SELECT
  company_users.display_name,
  company_users.user_id,
  company_users.company_id,
  company_users.id AS employee_id,
  users.email AS email,
  roles.name AS role
FROM
  company_users
  INNER JOIN users ON company_users.user_id = users.id
  INNER JOIN roles ON company_users.role_id = roles.id
WHERE
  company_id = $1 AND user_id = $2;

-- name: GetEmployeeByUserID :one
SELECT
  company_users.display_name,
  company_users.user_id,
  company_users.company_id,
  company_users.id AS employee_id,
  users.email AS email,
  roles.name AS role
FROM
  company_users
  INNER JOIN users ON company_users.user_id = users.id
  INNER JOIN roles ON company_users.role_id = roles.id
WHERE
  user_id = $1;