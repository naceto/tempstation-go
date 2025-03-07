-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  name, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
set 
  name = $2,
  email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListSensors :many
SELECT * FROM sensors
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetSensor :one
SELECT * FROM sensors
WHERE id = $1 LIMIT 1;

-- name: CreateSensor :one
INSERT INTO sensors (
  user_id, name, type, mac_address
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

