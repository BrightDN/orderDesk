-- name: GetCompanyEmployees :many
SELECT
  company_users.display_name,
  company_users.user_id,
  users.email AS email,
  roles.name AS role
FROM
  company_users
  INNER JOIN users ON company_users.user_id = users.id
  INNER JOIN roles ON company_users.role_id = roles.id
WHERE
  company_id = $1;