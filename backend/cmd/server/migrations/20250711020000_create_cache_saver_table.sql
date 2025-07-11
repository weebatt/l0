-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cache_saver (
    order_uid UUID PRIMARY KEY REFERENCES orders(order_uid) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cache_saver;
-- +goose StatementEnd
