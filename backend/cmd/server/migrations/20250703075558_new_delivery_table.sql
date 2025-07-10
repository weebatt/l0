-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE delivery (
    delivery_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_uid UUID UNIQUE,
    name VARCHAR(100),
    phone VARCHAR(20),
    zip VARCHAR(10),
    city VARCHAR(100),
    address VARCHAR(255),
    region VARCHAR(100),
    email VARCHAR(100),
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS delivery;
-- +goose StatementEnd
