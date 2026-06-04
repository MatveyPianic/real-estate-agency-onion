-- +goose Up
-- +goose StatementBegin

ALTER TABLE properties
    ADD COLUMN address VARCHAR(500),
    ADD COLUMN floor SMALLINT,
    ADD COLUMN total_floors SMALLINT,
    ADD COLUMN latitude DECIMAL(10,8),
    ADD COLUMN longitude DECIMAL(11,8);

ALTER TABLE properties
    ADD CONSTRAINT properties_floors_check
    CHECK (floor > 0 AND floor <= total_floors);

ALTER TABLE properties
    ADD CONSTRAINT properties_total_floors_check
    CHECK (total_floors > 0);

ALTER TABLE properties
    ADD CONSTRAINT properties_coordinates_check
    CHECK ((latitude IS NULL AND longitude IS NULL) OR 
    (latitude IS NOT NULL AND longitude IS NOT NULL));

CREATE INDEX IF NOT EXISTS ix_properties_coordinates ON properties(latitude, longitude) WHERE latitude IS NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE properties
    DROP CONSTRAINT IF EXISTS properties_floors_check,
    DROP CONSTRAINT IF EXISTS properties_total_floors_check,
    DROP CONSTRAINT IF EXISTS properties_coordinates_check;

ALTER TABLE properties
    DROP COLUMN IF EXISTS address,
    DROP COLUMN IF EXISTS floor,
    DROP COLUMN IF EXISTS total_floors,
    DROP COLUMN IF EXISTS latitude,
    DROP COLUMN IF EXISTS longitude;

DROP INDEX IF EXISTS ix_properties_coordinates;

-- +goose StatementEnd