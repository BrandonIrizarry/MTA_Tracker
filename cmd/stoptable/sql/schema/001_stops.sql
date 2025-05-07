-- +goose Up
CREATE TABLE stops(
       id INT PRIMARY KEY,
       name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE stops;
