-- name: CreateStop :one
INSERT INTO stops (
  id, name
) VALUES (
  ?, ?
)

RETURNING *;
