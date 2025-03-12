package sqlc

import (
	"context"
	"fmt"

	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/models"
)

func (s *storage) CreateSensor(ctx context.Context, createSensor *models.CreateSensor) (*models.Sensor, error) {
	params := s.convert.CreateSensorModelToDB(createSensor)
	cs, err := s.db.CreateSensor(ctx, params)
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	sensor := s.convert.SensorModelFromDB(cs)
	return sensor, nil
}

func (s *storage) GetSensor(ctx context.Context, id int64) (*models.Sensor, error) {
	sensor, err := s.db.GetSensor(ctx, id)
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	sensorModel := s.convert.SensorModelFromDB(sensor)
	return sensorModel, nil
}

func (s *storage) ListSensors(ctx context.Context, params *models.ListSensorsParams) ([]*models.Sensor, error) {
	var limit int32 = DefaultLimit
	if params.Limit != 0 {
		limit = params.Limit
	}

	sensors, err := s.db.ListSensors(ctx, db.ListSensorsParams{
		Limit:  limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	sensorModels := s.convert.SensorModelsFromDB(sensors)
	return sensorModels, nil
}
