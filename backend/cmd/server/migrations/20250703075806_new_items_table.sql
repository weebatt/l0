-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
       
CREATE TABLE items (
    chrt_id BIGINT PRIMARY KEY,
    order_uid UUID,
    track_number VARCHAR(50),
    price INT,
    rid VARCHAR(50),
    name VARCHAR(100),
    sale INT,
    size VARCHAR(10),
    total_price INT,
    nm_id BIGINT,
    brand VARCHAR(100),
    status INT,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
