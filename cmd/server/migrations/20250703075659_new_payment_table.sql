-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
       
CREATE TABLE payment (
     transaction UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
     order_uid UUID UNIQUE,
     request_id VARCHAR(50),
     currency VARCHAR(3),
     provider VARCHAR(50),
     amount INT,
     payment_dt BIGINT,
     bank VARCHAR(50),
     delivery_cost INT,
     goods_total INT,
     custom_fee INT,
     FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment;
-- +goose StatementEnd
