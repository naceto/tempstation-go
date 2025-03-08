package storage

import (
	"context"

	"github.com/naceto/tempstation/internal/storage/models"
)

type Storage interface {
	// Sensor
	CreateSensor(ctx context.Context, sensor *models.CreateSensor) (*models.Sensor, error)
	ListSensors(ctx context.Context, params *models.ListSensorsParams) ([]*models.Sensor, error)
	GetSensor(ctx context.Context, id int64) (*models.Sensor, error)

	// SensorData

	// User
	CreateUser(ctx context.Context, arg *models.CreateUser) (*models.User, error)
	ListUsers(ctx context.Context, params *models.ListUsersParams) ([]*models.User, error)
	GetUser(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.UpdateUser) (*models.User, error)
	DeleteUser(ctx context.Context, id int64) error
}
