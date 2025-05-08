-- +goose Up
CREATE TABLE stops(
       id TEXT PRIMARY KEY,
       name TEXT NOT NULL
);

-- +goose Down
DROP TABLE stops;
