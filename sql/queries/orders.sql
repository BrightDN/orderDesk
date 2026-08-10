-- name: CreateOrder :exec
INSERT INTO
  orders (id, supplier_id, employee_id, company_id)
VALUES
  ($1, $2, $3, $4);

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE
  id = $1;


-- name: GetOrdersBySupplier :many
SELECT
    orders.id,
    orders.created_at,
    employees.display_name AS placed_by,
    (
        SELECT count(*)
        FROM order_items
        WHERE order_id = orders.id
    ) AS item_count
FROM orders
INNER JOIN employees
    ON orders.employee_id = employees.id
WHERE
    orders.supplier_id = $1
    AND orders.company_id = $2;
