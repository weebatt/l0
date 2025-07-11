-- +goose Up
-- +goose StatementBegin
ALTER TABLE items
    ADD CONSTRAINT fk_items_order_uid
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
    ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items
    DROP CONSTRAINT IF EXISTS fk_items_order_uid;
-- +goose StatementEnd
