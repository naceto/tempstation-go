package resources

import (
	"context"
	"log/slog"

	api "github.com/naceto/tempstation/internal/generated/api/sensors"
	"github.com/naceto/tempstation/internal/resources/convert"
	"github.com/naceto/tempstation/internal/storage"
)

var _ api.StrictServerInterface = &Sensors{}

type Sensors struct {
	log     *slog.Logger
	convert convert.Convert
	store   storage.Storage
}

func NewSensors(log *slog.Logger, convert convert.Convert, store storage.Storage) *Sensors {
	return &Sensors{
		log:     log,
		convert: convert,
		store:   store,
	}
}

// (GET /v1/sensors)
func (s *Sensors) GetV1Sensors(ctx context.Context, request api.GetV1SensorsRequestObject) (api.GetV1SensorsResponseObject, error) {
	params := s.convert.ListSensorsAPIToStorage(request)
	sensors, err := s.store.ListSensors(ctx, params)
	if err != nil {
		s.log.Error("resources.GetV1Sensors", "store.ListSensors error", err)
		return nil, err
	}

	response := s.convert.ListSensorsAPIFromStorage(sensors)
	return response, nil
}

// (POST /v1/sensors)
func (s *Sensors) PostV1Sensors(ctx context.Context, request api.PostV1SensorsRequestObject) (api.PostV1SensorsResponseObject, error) {
	sensorModel := s.convert.PostSensorAPIToStorage(request)
	sensor, err := s.store.CreateSensor(ctx, sensorModel)
	if err != nil {
		s.log.Error("resources.PostV1Sensors", "store.CreateSensor error", err)
		return nil, err
	}

	response := s.convert.PostSensorAPIFromStorage(sensor)
	return response, nil
}

// (GET /v1/sensors/{id})
func (s *Sensors) GetV1SensorsId(ctx context.Context, request api.GetV1SensorsIdRequestObject) (api.GetV1SensorsIdResponseObject, error) {
	sensor, err := s.store.GetSensor(ctx, request.Id)
	if err != nil {
		s.log.Error("resources.GetV1SensorsId", "store.GetSensor error", err)
		return nil, err
	}

	response := s.convert.GetSensorAPIFromStorage(sensor)
	return response, nil
}

// (POST /v1/sensors/{id}/data)
func (s *Sensors) PostV1SensorsIdData(ctx context.Context, request api.PostV1SensorsIdDataRequestObject) (api.PostV1SensorsIdDataResponseObject, error) {
	sd := s.convert.PostSensorDataAPIToStorage(request)
	sensorData, err := s.store.CreateSensorData(ctx, sd)
	if err != nil {
		s.log.Error("resources.PostV1SensorsIdData", "store.CreateSensorData error", err)
		return nil, err
	}

	response := s.convert.SensorDataAPIFromStorage(sensorData)
	return response, nil
}
