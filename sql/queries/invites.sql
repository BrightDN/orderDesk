-- name: GetInvite :one
SELECT *
FROM invites
WHERE id = $1;

-- name: GetInvites :many
SELECT *
FROM invites;

-- name: GetCompanyInvites :many
SELECT
    invites.*,
    companies.name AS company_name
FROM invites
JOIN companies
    ON invites.company_id = companies.id
WHERE invites.invite_type = 'company';

-- name: GetEmployeeInvites :many
SELECT *
FROM invites
WHERE invite_type = "employee";

-- name: CreateInvite :exec
INSERT INTO invites (email, company_id, invite_type, token, expires_at)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
);

-- name: DeleteInvite :exec
DELETE FROM invites
WHERE id = $1;

-- name: DeleteUsedInvites :exec
DELETE FROM invites
WHERE used_at < now();

-- name: RenewInvite :exec
UPDATE invites
SET expires_at = $1,
    updated_at = now()
WHERE id = $2;