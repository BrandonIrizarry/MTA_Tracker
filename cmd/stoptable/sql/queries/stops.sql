-- name: CreateStop :exec
INSERT INTO stops (
  id, name
) VALUES (
  ?, ?
);

-- name: ClearAllStops :exec
DELETE FROM stops;
