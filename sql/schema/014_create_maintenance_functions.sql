-- +goose Up

-- This function is intentionally invoked by an application job or database scheduler;
-- PostgreSQL itself does not run scheduled jobs without an extension such as pg_cron.
-- +goose StatementBegin
CREATE FUNCTION purge_deleted_companies()
RETURNS TABLE (id INTEGER)
LANGUAGE sql
AS $$
    DELETE FROM companies
    WHERE deleted_at <= now() - INTERVAL '30 days'
    RETURNING companies.id;
$$;
-- +goose StatementEnd

-- Inserts a JSON array of order items atomically. Each item needs order_id,
-- product_id (or null), quantity, and name_at_order.
-- +goose StatementBegin
CREATE FUNCTION create_order_items(items JSONB)
RETURNS SETOF order_items
LANGUAGE plpgsql
AS $$
BEGIN
    IF jsonb_typeof(items) <> 'array' THEN
        RAISE EXCEPTION 'items must be a JSON array';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM jsonb_to_recordset(items) AS item(
            order_id TEXT,
            product_id INTEGER,
            quantity INTEGER,
            name_at_order TEXT
        )
        LEFT JOIN orders ON orders.id = item.order_id
        LEFT JOIN products ON products.id = item.product_id
        WHERE orders.id IS NULL
           OR item.quantity IS NULL
           OR item.quantity <= 0
           OR item.name_at_order IS NULL
           OR (item.product_id IS NOT NULL
               AND (products.id IS NULL OR products.supplier_id <> orders.supplier_id))
    ) THEN
        RAISE EXCEPTION 'an order item is invalid or does not belong to the order supplier';
    END IF;

    RETURN QUERY
    INSERT INTO order_items (order_id, product_id, quantity, name_at_order)
    SELECT item.order_id, item.product_id, item.quantity, item.name_at_order
    FROM jsonb_to_recordset(items) AS item(
        order_id TEXT,
        product_id INTEGER,
        quantity INTEGER,
        name_at_order TEXT
    )
    RETURNING order_items.*;
END;
$$;
-- +goose StatementEnd

-- +goose Down

DROP FUNCTION IF EXISTS create_order_items(JSONB);
DROP FUNCTION IF EXISTS purge_deleted_companies();
