-- +goose Up
-- +goose StatementBegin

-- админы
CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin

DROP TABLE IF EXISTS admins;

-- +goose StatementEnd