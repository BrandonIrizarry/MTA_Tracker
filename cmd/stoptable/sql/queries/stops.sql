-- name: CreateStop :exec
INSERT INTO stops (
  stop_id, name, route_id
) VALUES (
  ?, ?, ?
);

-- name: ClearAllStops :exec
DELETE FROM stops;

-- name: TestRouteExists :one
SELECT EXISTS (
       SELECT route_id FROM stops
       WHERE route_id = ?
);

-- name: ClearStopsByRoute :exec
DELETE FROM stops
WHERE route_id = ?;
