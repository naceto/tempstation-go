package models

import (
	"time"
)

type CreateUser struct {
	Name     string
	Email    string
	Password []byte
}

// User represents a user in the system.
type User struct {
	ID       int64
	Name     string
	Email    string
	Password []byte
}

type UpdateUser struct {
	ID   int64
	Name string
}

type UpdateUserPassword struct {
	ID          int64
	NewPassword []byte
}

// ListUser represents a user in the system.
type ListUser struct {
	ID    int64
	Name  string
	Email string
}

type SensorType string

const (
	SensorTypeDHT11 SensorType = "DHT11"
	SensorTypeDHT22 SensorType = "DHT22"
)

type Sensor struct {
	ID         int64
	UserID     int64
	Name       string
	Type       SensorType
	MacAddress string
}

type CreateSensor struct {
	UserID     int64
	Name       string
	Type       SensorType
	MacAddress string
}

type ListSensorsParams struct {
	Limit  int32
	Offset int32
	Mac    *string
}

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

// Sensor represents a sensor data in the system.
type SensorData struct {
	ID          int64
	SensorID    int64
	Temperature float32
	Humidity    float32
	ReadingTime time.Time
}

type CreateSensorData struct {
	SensorID    int64
	Temperature float32
	Humidity    float32
}

type GetSensorDataParams struct {
	SensorID int64
	Start    time.Time
	End      *time.Time
}
