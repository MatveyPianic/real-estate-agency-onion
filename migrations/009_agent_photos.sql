-- +goose Up
-- +goose StatementBegin

ALTER TABLE agents ADD COLUMN IF NOT EXISTS photo_path TEXT DEFAULT '';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE agents DROP COLUMN IF EXISTS photo_path;

-- +goose StatementEnd
