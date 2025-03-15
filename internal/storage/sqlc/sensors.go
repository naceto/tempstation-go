package sqlc

import (
	"context"
	"fmt"
	"time"

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

func (s *storage) CreateSensorData(ctx context.Context, data *models.CreateSensorData) (*models.SensorData, error) {
	dbData := s.convert.CreateSensorDataModelToDB(data)
	sd, err := s.db.CreateSensorData(ctx, dbData)
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	modelSensorData := s.convert.CreateSensorDataModelFromDB(sd)
	return modelSensorData, nil
}

func (s *storage) GetSensorData(ctx context.Context, params *models.GetSensorDataParams) ([]*models.SensorData, error) {
	if params.End == nil {
		now := time.Now()
		params.End = &now
	}

	p := s.convert.GetSensorDataParamsToDB(params)
	data, err := s.db.GetSensorData(ctx, p)
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	sensorDataList := s.convert.GetSensorDataFromDB(data)
	return sensorDataList, nil
}
