-- name: GetCompanySuppliers :many
SELECT
    suppliers.id,
    suppliers.name,
    suppliers.email,
    suppliers.contact,
    suppliers.deleted_at,
    COALESCE(order_mails.subject, 'Order') AS mail_subject,
    COALESCE(order_mails.mail_content, 'See order in attachment') AS mail_content,
    (
        SELECT COUNT(*)
        FROM products
        WHERE products.supplier_id = suppliers.id
    ) AS product_count
FROM suppliers
LEFT JOIN order_mails
    ON suppliers.id = order_mails.supplier_id
WHERE suppliers.company_id = $1
ORDER BY suppliers.created_at ASC;

-- name: GetCompanySupplier :one
SELECT
    suppliers.id,
    suppliers.name,
    suppliers.email,
    suppliers.contact,
    suppliers.deleted_at,
    COALESCE(order_mails.subject, 'Order') AS mail_subject,
    COALESCE(order_mails.mail_content, 'See order in attachment') AS mail_content,
    (
        SELECT COUNT(*)
        FROM products
        WHERE products.supplier_id = suppliers.id
    ) AS product_count
FROM suppliers
    LEFT JOIN order_mails
        ON suppliers.id = order_mails.supplier_id
    WHERE suppliers.name = $1 AND suppliers.company_id = $2;

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