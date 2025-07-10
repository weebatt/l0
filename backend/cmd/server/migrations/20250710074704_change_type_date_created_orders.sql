-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
    DROP COLUMN date_created,
    ADD COLUMN date_created VARCHAR(20) NOT NULL DEFAULT '2006-01-02T15:04:05Z';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
    DROP COLUMN date_created,
    ADD COLUMN date_created TIMESTAMP WITH TIME ZONE;
-- +goose StatementEnd
