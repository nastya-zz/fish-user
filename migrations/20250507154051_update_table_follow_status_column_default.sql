-- +goose Up
-- +goose StatementBegin
ALTER TABLE follows ALTER COLUMN status SET DEFAULT 'accepted';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
