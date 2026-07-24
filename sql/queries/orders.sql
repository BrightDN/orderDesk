-- Name: CreateOrder :one
INSERT INTO
  orders (id, supplier_id, employee_id)
VALUES
  ($1, $2, $3)
RETURNING
  $1;

-- Name: DeleteOrder :one
DELETE FROM orders
WHERE
  id = $1
RETURNING
  $1;