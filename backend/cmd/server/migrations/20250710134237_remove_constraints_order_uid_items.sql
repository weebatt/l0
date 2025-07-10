-- +goose Up
-- +goose StatementBegin
ALTER TABLE items
    DROP COLUMN order_uid,
    ADD COLUMN order_uid UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items
    DROP COLUMN order_uid,
    ADD COLUMN order_uid UUID UNIQUE;
-- +goose StatementEnd
