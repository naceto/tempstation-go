package resources

import (
	"encoding/json"
	"net/http"

	api "github.com/naceto/tempstation/internal/generated/api/sensors"
)

// optional code omitted
type SensorsResource struct{}

func NewSensorsResource() *SensorsResource {
	return &SensorsResource{}
}

// (GET /ping)
func (s *SensorsResource) GetV1Sensors(w http.ResponseWriter, r *http.Request) {
	resp := api.SensorsResponse{
		Sensors: []api.Sensor{
			{},
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
