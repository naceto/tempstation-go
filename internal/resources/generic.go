package resources

import (
	"net/http"
	"time"

	api "github.com/naceto/tempstation/internal/generated/api/generic"
	"github.com/naceto/tempstation/pkg/rest"
	"github.com/naceto/tempstation/pkg/sdk"
)

var _ api.ServerInterface = &Generic{}

type Generic struct{}

func NewGeneric() *Generic {
	return &Generic{}
}

// (GET /ping)
func (g *Generic) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := api.Pong{
		Ping:      "pong",
		Timestamp: sdk.Ptr(time.Now().UTC()),
	}

	rest.EncodeAndReturn(w, http.StatusOK, resp, nil)
}
