-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS inquiries(
    id BIGSERIAL PRIMARY KEY,
    property_id BIGINT NOT NULL REFERENCES properties(id) ON DELETE RESTRICT,
    client_name VARCHAR(255) NOT NULL,
    client_phone_number VARCHAR(100) NOT NULL,
    client_email VARCHAR(300),
    comment VARCHAR(1000),
    status VARCHAR(30) NOT NULL CHECK (status in ('new', 'scheduled', 'closed')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- индекс по недвижке
CREATE INDEX IF NOT EXISTS ix_inquiries_property_id ON inquiries(property_id);

-- индекс по cтатусам
CREATE INDEX IF NOT EXISTS ix_inquiries_status ON inquiries(status);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin

DROP TABLE IF EXISTS inquiries;

-- +goose StatementEnd