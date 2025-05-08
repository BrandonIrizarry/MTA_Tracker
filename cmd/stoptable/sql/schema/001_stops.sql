-- +goose Up
CREATE TABLE stops(
       id TEXT PRIMARY KEY,
       name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE stops;
