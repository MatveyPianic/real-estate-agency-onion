-- +goose Up
-- +goose StatementBegin

-- колонки 
ALTER TABLE properties ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE agents ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inquiries ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE showings ADD COLUMN deleted_at TIMESTAMPTZ;

-- индексы
CREATE INDEX IF NOT EXISTS ix_properties_deleted_at ON properties(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS ix_agents_deleted_at ON agents(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS ix_inquiries_deleted_at ON inquiries(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS ix_showings_deleted_at ON showings(id) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- удалить индексы
DROP INDEX IF EXISTS ix_properties_deleted_at;
DROP INDEX IF EXISTS ix_agents_deleted_at;
DROP INDEX IF EXISTS ix_inquiries_deleted_at;
DROP INDEX IF EXISTS ix_showings_deleted_at;

-- удалить колонки
ALTER TABLE properties DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE agents DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE inquiries DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE showings DROP COLUMN IF EXISTS deleted_at;

-- +goose StatementEnd
