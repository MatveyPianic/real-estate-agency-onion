-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS agents(
    id BIGSERIAL PRIMARY KEY,
    last_name VARCHAR(50) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    middle_name VARCHAR(50),
    phone TEXT NOT NULL,
    telegram VARCHAR(50),
    whatsapp VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- индекс чтобы не было дублей телефонов
CREATE UNIQUE INDEX IF NOT EXISTS ux_agents_phone on agents(phone);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS agents;

-- +goose StatementEnd