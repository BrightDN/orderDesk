-- +goose Up

CREATE INDEX order_mails_idx_supplier_id ON "order_mails" ("supplier_id");
CREATE INDEX products_idx_supplier_id ON "products" ("supplier_id");
CREATE INDEX suppliers_idx_company_id_id ON "suppliers" ("company_id","id");

-- +goose Down

DROP INDEX IF EXISTS order_mails_idx_supplier_id;
DROP INDEX IF EXISTS products_idx_supplier_id;
DROP INDEX IF EXISTS suppliers_idx_company_id_id;