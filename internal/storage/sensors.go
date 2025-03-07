package storage

import (
	"context"

	"github.com/naceto/tempstation/internal/generated/db"
)

func (s *storage) CreateSensor(ctx context.Context, arg db.CreateSensorParams) (db.Sensor, error) {
	return s.db.CreateSensor(ctx, arg)
}

func (s *storage) GetSensor(ctx context.Context, id int64) (db.Sensor, error) {
	return s.db.GetSensor(ctx, id)
}

func (s *storage) ListSensors(ctx context.Context) ([]db.Sensor, error) {
	return s.db.ListSensors(ctx, db.ListSensorsParams{})
}
