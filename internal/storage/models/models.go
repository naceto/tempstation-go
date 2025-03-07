package models

import (
	"database/sql"
)

// User represents a user in the system.
type User struct {
	ID    int64          `json:"id"`
	Name  sql.NullString `json:"name"`
	Email sql.NullString `json:"email"`
}

// Sensor represents a sensor in the system.
type Sensor struct {
	ID          int64          `json:"id"`
	SensorID    sql.NullInt64  `json:"sensor_id"`
	Temperature sql.NullString `json:"temperature"`
	Humidity    sql.NullString `json:"humidity"`
	ReadingTime sql.NullTime   `json:"reading_time"`
	MacAddress  string         `json:"mac_address"`
}
