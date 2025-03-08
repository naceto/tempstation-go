package resources

import (
	"context"
	"log/slog"

	api "github.com/naceto/tempstation/internal/generated/api/users"
	"github.com/naceto/tempstation/internal/storage"
)

var _ api.StrictServerInterface = &Users{}

type Users struct {
	log   *slog.Logger
	store storage.Storage
}

func NewUsers(log *slog.Logger, store storage.Storage) *Users {
	return &Users{
		log:   log,
		store: store,
	}
}

// (GET /v1/users)
func (u *Users) GetV1Users(ctx context.Context, request api.GetV1UsersRequestObject) (api.GetV1UsersResponseObject, error) {
	return nil, nil
}

// (GET /v1/users/{id})
func (u *Users) GetV1UsersId(ctx context.Context, request api.GetV1UsersIdRequestObject) (api.GetV1UsersIdResponseObject, error) {
	return nil, nil
}

// (POST /v1/users)
func (u *Users) PostV1Users(ctx context.Context, request api.PostV1UsersRequestObject) (api.PostV1UsersResponseObject, error) {
	return nil, nil
}
