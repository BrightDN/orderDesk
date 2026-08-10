-- name: PurgeDeletedCompanies :many
SELECT purge_deleted_companies AS id FROM purge_deleted_companies();

-- name: CreateOrderItems :many
SELECT * FROM create_order_items($1::jsonb);
