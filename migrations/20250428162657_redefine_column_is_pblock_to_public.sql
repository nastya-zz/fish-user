-- +goose Up
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN is_blocked TO is_public;
ALTER TABLE users ALTER COLUMN is_public SET DEFAULT TRUE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
