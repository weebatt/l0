-- +goose Up
-- +goose StatementBegin
ALTER TABLE items
    DROP COLUMN transaction,
    ADD COLUMN order_uid UUID UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items
    DROP COLUMN order_uid,
    ADD COLUMN transaction UUID UNIQUE;
-- +goose StatementEnd
