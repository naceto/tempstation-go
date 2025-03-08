package sqlc

import (
	"context"

	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/models"
)

func (s *storage) CreateSensor(ctx context.Context, sensor *models.CreateSensor) (*models.Sensor, error) {
	cs, err := s.db.CreateSensor(ctx, db.CreateSensorParams{
		UserID:     sensor.UserID,
		Name:       sensor.Name,
		Type:       db.SensorType(sensor.Type),
		MacAddress: sensor.MacAddress,
	})
	if err != nil {
		return nil, err
	}

	return &models.Sensor{
		ID:         cs.ID,
		UserID:     cs.UserID,
		Name:       cs.Name,
		Type:       models.SensorType(cs.Type),
		MacAddress: cs.MacAddress,
	}, nil
}

func (s *storage) GetSensor(ctx context.Context, id int64) (*models.Sensor, error) {
	sensor, err := s.db.GetSensor(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Sensor{
		ID:         sensor.ID,
		UserID:     sensor.UserID,
		Name:       sensor.Name,
		Type:       models.SensorType(sensor.Type),
		MacAddress: sensor.MacAddress,
	}, nil
}

func (s *storage) ListSensors(ctx context.Context, params *models.ListSensorsParams) ([]*models.Sensor, error) {
	sensors, err := s.db.ListSensors(ctx, db.ListSensorsParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}

	results := []*models.Sensor{}
	for _, sen := range sensors {
		results = append(results, &models.Sensor{
			ID:         sen.ID,
			UserID:     sen.UserID,
			Name:       sen.Name,
			Type:       models.SensorType(sen.Type),
			MacAddress: sen.MacAddress,
		})
	}

	return results, nil
}
