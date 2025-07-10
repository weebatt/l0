-- +goose Up
-- +goose StatementBegin
ALTER TABLE delivery
    DROP COLUMN delivery_id,
    ADD COLUMN delivery_id SERIAL PRIMARY KEY;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE delivery
    DROP COLUMN delivery_id,
    ADD COLUMN delivery_id UUID PRIMARY KEY DEFAULT uuid_generate_v4();
-- +goose StatementEnd
