-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
    ALTER COLUMN order_uid DROP DEFAULT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
    ALTER COLUMN order_uid SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd
