-- +goose Up

CREATE TABLE IF NOT EXISTS feed(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX feed_user_idx ON feed (user_id);

-- +goose Down

DROP INDEX IF EXISTS feed_user_idx;
DROP TABLE IF EXISTS feed;