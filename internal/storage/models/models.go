package models

import (
	"time"
)

type CreateUser struct {
	Name  string
	Email string
}

// User represents a user in the system.
type User struct {
	ID    int64
	Name  string
	Email string
}

type UpdateUser struct {
	Name string
}

type SensorType string

const (
	SensorTypeDht11 SensorType = "dht11"
	SensorTypeDht22 SensorType = "dht22"
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
	MacAddress  string
}

type CreateSensorData struct {
	SensorID    int64
	Temperature float32
	Humidity    float32
	ReadingTime time.Time
	MacAddress  string
}
