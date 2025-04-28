-- +goose Up
-- +goose StatementBegin
ALTER TABLE users DROP CONSTRAINT users_username_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
