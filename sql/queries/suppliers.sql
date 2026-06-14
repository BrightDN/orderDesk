-- name: GetCompanySuppliers :many
SELECT
    suppliers.id,
    suppliers.name,
    suppliers.email,
    suppliers.contact,
    suppliers.deleted_at,
    COUNT(products.id) AS product_count
FROM
    suppliers
    LEFT JOIN products ON suppliers.id = products.supplier_id
WHERE
    suppliers.company_id = $1
GROUP BY
    suppliers.id;

-- name: CreateSupplier :one
INSERT INTO suppliers (name, email, company_id, contact)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSupplierByID :one
SELECT *
FROM suppliers
WHERE id = $1;

-- name: GetSupplierByNameAndCompany :one
SELECT *
FROM suppliers
WHERE name = $1 AND company_id = $2;

-- name: EditSupplierByNameAndCompanyID :exec
UPDATE suppliers 
    SET
        name = $3,
        email = $4,
        contact = $5,
        deleted_at = $6 
    WHERE
        name = $1 
        AND company_id = $2;

-- name: GetSupplierByName :one
SELECT *
FROM suppliers
WHERE name = $1;