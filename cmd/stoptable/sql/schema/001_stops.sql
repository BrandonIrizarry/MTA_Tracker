-- +goose Up
CREATE TABLE stops(
       stop_id TEXT NOT NULL,
       name TEXT NOT NULL,
       route_id TEXT NOT NULL,
       UNIQUE (stop_id, route_id)
);

-- +goose Down
DROP TABLE stops;
