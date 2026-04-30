-- +goose Up
ALTER TABLE users
ADD COLUMN session TEXT;

-- +goose Down
ALTER TABLE users
DROP COLUMN session;