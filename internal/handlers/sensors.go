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

// (GET /ping)
func (s *Sensors) GetV1Sensors(ctx context.Context, request api.GetV1SensorsRequestObject) (api.GetV1SensorsResponseObject, error) {
	resp := api.GetV1Sensors200JSONResponse{
		SensorsResponseJSONResponse: api.SensorsResponseJSONResponse{
			Sensors: []api.Sensor{
				{},
			},
		},
	}

	return resp, nil
}
