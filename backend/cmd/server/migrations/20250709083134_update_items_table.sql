-- +goose Up
-- +goose StatementBegin
ALTER TABLE items
    DROP COLUMN chrt_id,
    DROP COLUMN order_uid,
    ADD COLUMN chrt_id SERIAL PRIMARY KEY,
    ADD COLUMN order_uid UUID UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items
    DROP COLUMN chrt_id,
    DROP COLUMN order_uid,
    ADD COLUMN order_uid UUID,
    ADD COLUMN chrt_id BIGINT PRIMARY KEY;
-- +goose StatementEnd
