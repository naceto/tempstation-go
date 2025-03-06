package resources

import (
	"context"
	"log/slog"

	api "github.com/naceto/tempstation/internal/generated/api/sensors"
	"github.com/naceto/tempstation/internal/storage"
)

var _ api.StrictServerInterface = &Sensors{}

type Sensors struct {
	log   *slog.Logger
	store storage.Storage
}

func NewSensors(log *slog.Logger, store storage.Storage) *Sensors {
	return &Sensors{
		log:   log,
		store: store,
	}
}

// (GET /v1/sensors)
func (s *Sensors) GetV1Sensors(ctx context.Context, request api.GetV1SensorsRequestObject) (api.GetV1SensorsResponseObject, error) {
	// TODO: use goverter generated converters to convert from internal models to api models
	resp := api.GetV1Sensors200JSONResponse{
		SensorsDataResponseJSONResponse: api.SensorsDataResponseJSONResponse{
			Sensors: []api.SensorData{
				{},
			},
		},
	}

	return resp, nil
}

// (POST /v1/sensors)
func (s *Sensors) PostV1Sensors(ctx context.Context, request api.PostV1SensorsRequestObject) (api.PostV1SensorsResponseObject, error) {
	// TODO: use goverter generated converters to convert from internal models to api models
	return api.PostV1Sensors200JSONResponse{
		SensorResponseJSONResponse: api.SensorResponseJSONResponse{
			Id:   request.Body.UserId,
			Mac:  request.Body.Mac,
			Name: request.Body.Name,
			Type: request.Body.Type,
		},
	}, nil
}
