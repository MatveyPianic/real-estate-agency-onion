-- +goose Up
-- +goose StatementBegin

-- сессии
CREATE TABLE IF NOT EXISTS admin_sessions (
    id SERIAL PRIMARY KEY,
    admin_id INT NOT NULL REFERENCES admins(id) ON DELETE CASCADE,
    token TEXT UNIQUE NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- индекс по admin_id
CREATE INDEX IF NOT EXISTS ix_admin_sessions_admin_id ON admin_sessions(admin_id);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin

DROP TABLE IF EXISTS admin_sessions;

-- +goose StatementEnd