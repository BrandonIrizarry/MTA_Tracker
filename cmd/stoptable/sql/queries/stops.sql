-- name: CreateStop :exec
INSERT INTO stops (
  id, name, route_id
) VALUES (
  ?, ?, ?
);

-- name: ClearAllStops :exec
DELETE FROM stops;
