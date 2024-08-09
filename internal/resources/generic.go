package resources

import (
	"encoding/json"
	"net/http"

	api "github.com/naceto/tempstation/internal/generated/api/generic"
)

// optional code omitted
type GenericResource struct{}

func NewGenericResource() *GenericResource {
	return &GenericResource{}
}

// (GET /ping)
func (res *GenericResource) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := api.Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
