-- +goose Up
CREATE TABLE users (
  ID UUID PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
