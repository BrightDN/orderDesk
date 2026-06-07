-- name: GetProducts :many
SELECT *
FROM products
WHERE supplier_id = $1;

-- name: CreateProduct :exec
INSERT INTO products (name, supplier_id)
VALUES ($1, $2);

-- name: DeleteProduct :exec
DELETE FROM products
WHERE supplier_id = $1 AND id = $2;