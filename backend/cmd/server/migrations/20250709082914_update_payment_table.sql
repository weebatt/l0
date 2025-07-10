-- +goose Up
-- +goose StatementBegin
ALTER TABLE payment
    DROP CONSTRAINT IF EXISTS payment_order_uid_fkey,
    DROP COLUMN transaction,
    ADD COLUMN payment_id SERIAL PRIMARY KEY,
    ADD COLUMN transaction UUID UNIQUE,
    ADD CONSTRAINT fk_transaction_orders FOREIGN KEY (transaction) REFERENCES orders(order_uid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE payment
    DROP CONSTRAINT IF EXISTS fk_transaction_orders,
    DROP COLUMN payment_id,
    DROP COLUMN transaction,
    ADD COLUMN transaction UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ADD CONSTRAINT payment_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES orders(order_uid);
-- +goose StatementEnd
