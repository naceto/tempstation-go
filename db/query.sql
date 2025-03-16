-- name: CreateUser :one
INSERT INTO users (
  name, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  name = $2,
  email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateSensor :one
INSERT INTO sensors (
  user_id, name, type, mac_address
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateSensor :one
UPDATE sensors
SET
  name = $2,
  type = $3
WHERE id = $1
RETURNING *;

-- name: ListSensors :many
SELECT * FROM sensors
WHERE (mac_address = sqlc.narg('mac') OR sqlc.narg('mac') IS NULL)
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetSensor :one
SELECT * FROM sensors
WHERE id = $1 LIMIT 1;

-- name: GetSensorByMAC :one
SELECT * FROM sensors
WHERE mac_address = $1 LIMIT 1;

-- name: CreateSensorData :one
INSERT INTO sensor_data (
  sensor_id, temperature,
  humidity
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetSensorData :many
SELECT * FROM sensor_data
WHERE sensor_id = @id AND reading_time BETWEEN @start AND sqlc.arg('end');

