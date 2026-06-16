-- name: CreateOrderMail :exec
    INSERT INTO order_mails(supplier_id, subject, mail_content)
    VALUES (
      $1,
      'Order',
      'See order in attachment'
    );

-- name: UpdateOrderMail :one
UPDATE order_mails
SET subject = $1, mail_content = $2
WHERE supplier_id = $3
RETURNING *;
