package storage

import (
	"context"

	"github.com/naceto/tempstation/internal/storage/models"
)

type Storage interface {
	// User
	CreateUser(ctx context.Context, user *models.CreateUser) (*models.User, error)
	ListUsers(ctx context.Context, params *models.ListUsersParams) ([]*models.ListUser, error)
	GetUser(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.UpdateUser) (*models.User, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUserPassword(ctx context.Context, userPassword *models.UpdateUserPassword) error

	// Sensor
	CreateSensor(ctx context.Context, sensor *models.CreateSensor) (*models.Sensor, error)
	ListSensors(ctx context.Context, params *models.ListSensorsParams) ([]*models.Sensor, error)
	GetSensor(ctx context.Context, id int64) (*models.Sensor, error)

	// SensorData
	CreateSensorData(ctx context.Context, sensorData *models.CreateSensorData) (*models.SensorData, error)
	GetSensorData(ctx context.Context, params *models.GetSensorDataParams) ([]*models.SensorData, error)
}
