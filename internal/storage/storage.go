package storage

import (
	"context"

	"github.com/naceto/tempstation/internal/generated/db"
	// Import the validation package
)

type Storage interface {
	CreateSensor(ctx context.Context, arg db.CreateSensorParams) (db.Sensor, error)
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetSensor(ctx context.Context, id int64) (db.Sensor, error)
	GetUser(ctx context.Context, id int64) (db.User, error)
	ListSensors(ctx context.Context) ([]db.Sensor, error)
	ListUsers(ctx context.Context) ([]db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) error
}

type storage struct {
	db db.Querier
}

func NewStorage(db db.Querier) Storage {
	return &storage{
		db: db,
	}
}

// Force inteface implementation
var _ Storage = (*storage)(nil)
