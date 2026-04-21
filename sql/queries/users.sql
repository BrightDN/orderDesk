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