-- name: GetOrderMailData :one
SELECT
        order_mails.subject,
        order_mails.mail_content 
    FROM
        order_mails 
    WHERE
        order_mails.supplier_id = $1;