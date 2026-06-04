-- +goose Up
-- +goose StatementBegin

-- Добавляем новые колонки (сначала как NULL или с DEFAULT)
ALTER TABLE properties
    ADD COLUMN title VARCHAR(255),
    ADD COLUMN description TEXT,
    ADD COLUMN property_type VARCHAR(20) NOT NULL DEFAULT 'apartment',
    ADD COLUMN deal_type VARCHAR(10) NOT NULL DEFAULT 'sale',
    ADD COLUMN currency VARCHAR(3) NOT NULL DEFAULT 'USD';

-- check constraint для property_type
ALTER TABLE properties
    ADD CONSTRAINT properties_property_type_check
    CHECK (property_type IN ('apartment', 'house', 'garage', 'land', 'commercial'));

-- check constraint для deal_type
ALTER TABLE properties
    ADD CONSTRAINT properties_deal_type_check
    CHECK (deal_type IN ('sale', 'rent'));

-- Заполняем существующие строки
UPDATE properties SET title = 'Объект недвижимости' WHERE title IS NULL;
UPDATE properties SET description = 'Описание отсутствует' WHERE description IS NULL;

-- Теперь делаем обязательными
ALTER TABLE properties ALTER COLUMN title SET NOT NULL;
ALTER TABLE properties ALTER COLUMN description SET NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE properties
    DROP CONSTRAINT IF EXISTS properties_property_type_check,
    DROP CONSTRAINT IF EXISTS properties_deal_type_check;

ALTER TABLE properties
    DROP COLUMN IF EXISTS title,
    DROP COLUMN IF EXISTS description,
    DROP COLUMN IF EXISTS property_type,
    DROP COLUMN IF EXISTS deal_type,
    DROP COLUMN IF EXISTS currency;

-- +goose StatementEnd