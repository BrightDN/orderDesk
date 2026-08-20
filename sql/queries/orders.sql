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

-- name: GetOrderForDownload :many
SELECT
    orders.id,
    orders.created_at,
    suppliers.name AS supplier_name,
    suppliers.email AS supplier_email,
    companies.name AS company_name,
    employees.display_name AS placed_by,
    order_items.name_at_order,
    order_items.quantity
FROM orders
INNER JOIN suppliers
    ON orders.supplier_id = suppliers.id
INNER JOIN companies
    ON orders.company_id = companies.id
INNER JOIN employees
    ON orders.employee_id = employees.id
INNER JOIN order_items
    ON order_items.order_id = orders.id
WHERE
    orders.id = $1
    AND orders.company_id = $2
ORDER BY order_items.id;
