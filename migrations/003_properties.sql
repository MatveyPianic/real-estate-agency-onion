-- +goose Up
-- +goose StatementBegin

-- недвижимость
CREATE TABLE IF NOT EXISTS properties(
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE RESTRICT,
    district_id BIGINT NOT NULL REFERENCES districts(id) ON DELETE RESTRICT,
    price INTEGER NOT NULL  CHECK (price > 0),
    area INTEGER NOT NULL  CHECK (area > 0),
    rooms SMALLINT NOT NULL ,
    condition VARCHAR(30) NOT NULL CHECK (condition in ('new', 'good', 'needs_repair')),
    status VARCHAR(30) NOT NULL CHECK (status in ('draft', 'published', 'archived')), --черновик (не видно на сайте), опубликован, скрыт/снят
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- индекс на недвижимость
CREATE INDEX IF NOT EXISTS ix_property_district_id ON properties(district_id);

-- индекс на цену
CREATE INDEX IF NOT EXISTS ix_property_price ON properties(price);


-- индекс на площадь
CREATE INDEX IF NOT EXISTS ix_property_area ON properties(area);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin

DROP TABLE IF EXISTS properties;

-- +goose StatementEnd