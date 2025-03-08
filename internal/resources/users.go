package resources

import (
	"context"
	"log/slog"

	api "github.com/naceto/tempstation/internal/generated/api/users"
	"github.com/naceto/tempstation/internal/resources/convert"
	"github.com/naceto/tempstation/internal/resources/validate"
	"github.com/naceto/tempstation/internal/storage"
)

var _ api.StrictServerInterface = &Users{}

type Users struct {
	log     *slog.Logger
	convert convert.Convert
	store   storage.Storage
}

func NewUsers(log *slog.Logger, convert convert.Convert, store storage.Storage) *Users {
	return &Users{
		log:     log,
		convert: convert,
		store:   store,
	}
}

// (GET /v1/users)
func (u *Users) GetV1Users(ctx context.Context, request api.GetV1UsersRequestObject) (api.GetV1UsersResponseObject, error) {
	return nil, nil
}

// (GET /v1/users/{id})
func (u *Users) GetV1UsersId(ctx context.Context, request api.GetV1UsersIdRequestObject) (api.GetV1UsersIdResponseObject, error) {
	user, err := u.store.GetUser(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return u.convert.GetUserAPIFromStorage(user), nil
}

// (POST /v1/users)
func (u *Users) PostV1Users(ctx context.Context, request api.PostV1UsersRequestObject) (api.PostV1UsersResponseObject, error) {
	createUser := u.convert.CreateUserAPIToStorage(&request)
	if err := validate.CreateUser(createUser); err != nil {
		return nil, err
	}

	user, err := u.store.CreateUser(ctx, createUser)
	if err != nil {
		return nil, err
	}

	return u.convert.PostUserAPIFromStorage(user), nil
}
