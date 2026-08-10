-- +goose Up

-- These changes make the corrected constraints available to databases that
-- applied migrations 011–013 before those migrations were revised.
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conrelid = 'order_items'::regclass
          AND contype = 'c'
          AND pg_get_constraintdef(oid) = 'CHECK ((quantity > 0))'
    ) THEN
        ALTER TABLE order_items
            ADD CONSTRAINT order_items_quantity_positive CHECK (quantity > 0);
    END IF;
END;
$$;
-- +goose StatementEnd

ALTER TABLE invites
    DROP CONSTRAINT IF EXISTS invites_company_id_fkey;

ALTER TABLE invites
    ADD CONSTRAINT invites_company_id_fkey
    FOREIGN KEY (company_id)
    REFERENCES companies(id)
    ON DELETE CASCADE;

-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conrelid = 'order_mails'::regclass
          AND conname = 'order_mails_supplier_unique'
    ) THEN
        ALTER TABLE order_mails
            ADD CONSTRAINT order_mails_supplier_unique UNIQUE (supplier_id);
    END IF;
END;
$$;
-- +goose StatementEnd

-- +goose Down

ALTER TABLE order_mails
    DROP CONSTRAINT IF EXISTS order_mails_supplier_unique;

ALTER TABLE invites
    DROP CONSTRAINT IF EXISTS invites_company_id_fkey;

ALTER TABLE invites
    ADD CONSTRAINT invites_company_id_fkey
    FOREIGN KEY (company_id)
    REFERENCES companies(id)
    ON DELETE RESTRICT;

ALTER TABLE order_items
    DROP CONSTRAINT IF EXISTS order_items_quantity_positive;
