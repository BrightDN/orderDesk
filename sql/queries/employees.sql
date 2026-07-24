-- name: GetCompanyEmployees :many
SELECT
  employees.display_name,
  employees.user_id,
  employees.company_id,
  employees.id AS employee_id,
  users.email AS email,
  roles.name AS role
FROM
  employees
  INNER JOIN users ON employees.user_id = users.id
  INNER JOIN roles ON employees.role_id = roles.id
WHERE
  company_id = $1;

-- name: CreateCompanyEmployee :one
INSERT INTO employees (company_id, user_id, role_id, display_name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEmployee :one
SELECT
        employees.display_name,
        employees.user_id,
        employees.company_id,
        employees.id AS employee_id,
        companies.name AS employed_at,
        users.email AS email,
        roles.name AS role 
    FROM
        employees 
    INNER JOIN
        users 
            ON employees.user_id = users.id 
    INNER JOIN
        roles 
            ON employees.role_id = roles.id 
    INNER JOIN
        companies 
            ON companies.id = $3 
    WHERE
        employees.company_id = $1 
        AND employees.user_id = $2;

-- name: GetEmployeeByUserID :one
SELECT
  employees.display_name,
  employees.user_id,
  employees.company_id,
  employees.id AS employee_id,
  users.email AS email,
  roles.name AS role,
  companies.name AS employed_at
FROM
  employees
  INNER JOIN users ON employees.user_id = users.id
  INNER JOIN roles ON employees.role_id = roles.id
  INNER JOIN companies ON employees.company_id = companies.id
WHERE
  user_id = $1;