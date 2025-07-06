-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
       
CREATE TABLE IF NOT EXISTS orders (
    order_uid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    track_number VARCHAR(50) UNIQUE,
    entry VARCHAR(10),
    locale VARCHAR(10),
    customer_id VARCHAR(50),
    delivery_service VARCHAR(50),
    shardkey VARCHAR(10),
    sm_id INT,
    date_created TIMESTAMP WITH TIME ZONE,
    oof_shard VARCHAR(10),
    internal_signature VARCHAR(50)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
