-- +goose Up
-- +goose StatementBegin

-- таблица картинок
CREATE TABLE IF NOT EXISTS property_photos(
    id BIGSERIAL PRIMARY KEY,
    property_id BIGINT NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
    file_path VARCHAR(255) NOT NULL,
    is_cover BOOLEAN NOT NULL DEFAULT false
);

-- у объекта может быть только одна обложка
CREATE UNIQUE INDEX IF NOT EXISTS uix_property_photos_cover_per_property ON property_photos(property_id) WHERE is_cover = true;

-- индекс по property id, ускорит выборку из галлереи
CREATE INDEX IF NOT EXISTS ix_property_photos_property_id ON property_photos(property_id);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin

DROP TABLE IF EXISTS property_photos;

-- +goose StatementEnd