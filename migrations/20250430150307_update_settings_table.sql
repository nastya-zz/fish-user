-- +goose Up
-- +goose StatementBegin
ALTER TABLE settings RENAME COLUMN id TO user_id;
ALTER TABLE settings ALTER COLUMN user_id drop default;
ALTER TABLE settings ALTER COLUMN user_id SET NOT NULL;
drop table user_settings;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
