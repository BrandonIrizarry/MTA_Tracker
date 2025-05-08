-- name: CreateStop :exec
INSERT INTO stops (
  id, name
) VALUES (
  ?, ?
);
