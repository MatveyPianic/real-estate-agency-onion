-- +goose Up
-- +goose StatementBegin

-- показы

CREATE TABLE IF NOT EXISTS showings (
    id BIGSERIAL PRIMARY KEY,
    property_id BIGINT NOT NULL REFERENCES properties(id) ON DELETE RESTRICT,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE RESTRICT,
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ NOT NULL CHECK (ends_at > starts_at),
    status VARCHAR(30) NOT NULL CHECK (status in ('scheduled', 'done', 'canceled')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- индекс по агенту и времени начала
CREATE INDEX IF NOT EXISTS ix_showings_agent_id_starts_at ON showings(agent_id, starts_at);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin

DROP TABLE IF EXISTS showings;

-- +goose StatementEnd