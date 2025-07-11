-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
    ALTER COLUMN entry TYPE VARCHAR(20),
    ALTER COLUMN locale TYPE VARCHAR(20),
    ALTER COLUMN shardkey TYPE VARCHAR(20),
    ALTER COLUMN oof_shard TYPE VARCHAR(20),
    ALTER COLUMN internal_signature TYPE VARCHAR(100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
    ALTER COLUMN entry TYPE VARCHAR(10),
    ALTER COLUMN locale TYPE VARCHAR(10),
    ALTER COLUMN shardkey TYPE VARCHAR(10),
    ALTER COLUMN oof_shard TYPE VARCHAR(10),
    ALTER COLUMN internal_signature TYPE VARCHAR(10);
-- +goose StatementEnd
