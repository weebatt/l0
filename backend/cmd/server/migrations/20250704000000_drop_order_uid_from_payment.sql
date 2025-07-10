-- +goose Up
ALTER TABLE payment DROP COLUMN IF EXISTS order_uid;

-- +goose Down
ALTER TABLE payment ADD COLUMN order_uid UUID UNIQUE;
-- Optionally, re-add the foreign key if needed:
-- ALTER TABLE payment ADD CONSTRAINT payment_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES orders(order_uid);
