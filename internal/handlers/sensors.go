package handlers

import (
	"net/http"

	api "github.com/naceto/tempstation/internal/generated/api/sensors"
	"github.com/naceto/tempstation/pkg/rest"
)

var _ api.ServerInterface = &SensorsResource{}

type SensorsResource struct{}

func NewSensorsResource() *SensorsResource {
	return &SensorsResource{}
}

// (GET /ping)
func (s *SensorsResource) GetV1Sensors(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	resp := api.SensorsResponse{
		Sensors: []api.Sensor{
			{},
		},
	}

	rest.EncodeAndReturn(w, 200, resp, nil)
}
