-- name: GetUsers :many

SELECT *
FROM users;

-- name: GetUserByMail :one

SELECT *
FROM users
WHERE email = $1;

-- name: GetUserById :one

SELECT *
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO
  users (email, password, name, is_admin)
VALUES
  ($1, $2, $3, $4)
RETURNING
  *;

-- name: GetCompanyCount :one
SELECT COUNT(*)
FROM company_users
WHERE user_id = $1;