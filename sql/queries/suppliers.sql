-- name: GetCompanySuppliers :many
SELECT
    suppliers.id,
    suppliers.name,
    suppliers.email,
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
