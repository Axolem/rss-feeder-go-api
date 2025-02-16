-- +goose Up

ALTER TABLE users ADD COLUMN IF NOT EXISTS api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
  encode(sha256(random()::text::bytea), 'hex')
);

CREATE INDEX user_api_key_id ON users (id) INCLUDE (name, api_key, created_at, updated_at);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS user_api_key_id;
ALTER TABLE users DROP COLUMN api_key;