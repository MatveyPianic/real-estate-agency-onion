-- +goose Up
-- +goose StatementBegin

-- города
CREATE TABLE IF NOT EXISTS cities (
        id          BIGSERIAL PRIMARY KEY,
        name        TEXT NOT NULL,
        created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
    );

-- чтобы не было дублей городов с одинаковым названием
CREATE UNIQUE INDEX IF NOT EXISTS ux_cities_name ON cities (name);

-- районы
CREATE TABLE IF NOT EXISTS districts (
        id          BIGSERIAL PRIMARY KEY,
        city_id     BIGINT NOT NULL REFERENCES cities(id) ON DELETE RESTRICT,
        name        TEXT NOT NULL,
        created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
    );

-- в одном городе районы не должны повторяться
CREATE UNIQUE INDEX IF NOT EXISTS ux_districts_city_id_name ON districts (city_id, name);

-- индекс ускорит фильтр "все районы города" и join-ы
CREATE INDEX IF NOT EXISTS ix_districts_city_id ON districts (city_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS districts;
DROP TABLE IF EXISTS cities;

-- +goose StatementEnd