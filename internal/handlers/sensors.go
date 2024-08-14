package handlers

import (
	"context"

	api "github.com/naceto/tempstation/internal/generated/api/sensors"
)

var _ api.StrictServerInterface = &Sensors{}

type Sensors struct{}

func NewSensors() *Sensors {
	return &Sensors{}
}

// (GET /v1/sensors)
func (s *Sensors) GetV1Sensors(ctx context.Context, request api.GetV1SensorsRequestObject) (api.GetV1SensorsResponseObject, error) {
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
	return api.PostV1Sensors200JSONResponse{
		SensorResponseJSONResponse: api.SensorResponseJSONResponse{
			Id:   request.Body.UserId,
			Mac:  request.Body.Mac,
			Name: request.Body.Name,
			Type: request.Body.Type,
		},
	}, nil
}
