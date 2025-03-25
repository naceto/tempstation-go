package resources

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	api "github.com/naceto/tempstation/internal/generated/api/users"
	"github.com/naceto/tempstation/internal/resources/convert"
	"github.com/naceto/tempstation/internal/resources/validate"
	"github.com/naceto/tempstation/internal/storage"
	"github.com/naceto/tempstation/internal/storage/models"
	"golang.org/x/crypto/bcrypt"
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
	params := u.convert.ListUsersAPIToStorage(request)
	users, err := u.store.ListUsers(ctx, params)
	if err != nil {
		u.log.Error("resources.GetV1Users", "store.ListUser error", err)
		return nil, err
	}

	response := u.convert.ListUsersAPIFromStorage(users)
	return response, nil
}

// (GET /v1/users/{id})
func (u *Users) GetV1UsersId(ctx context.Context, request api.GetV1UsersIdRequestObject) (api.GetV1UsersIdResponseObject, error) {
	user, err := u.store.GetUser(ctx, request.Id)
	if err != nil {
		u.log.Error("resources.GetV1UsersId", "store.GetUser error", err)
		return nil, err
	}

	response := u.convert.GetUserAPIFromStorage(user)
	return response, nil
}

// (POST /v1/users)
func (u *Users) PostV1Users(ctx context.Context, request api.PostV1UsersRequestObject) (api.PostV1UsersResponseObject, error) {
	createUser := u.convert.PostUserAPIToStorage(&request)

	if err := validate.CreateUser(createUser); err != nil {
		u.log.Error("resources.PostV1Users", "validate.CreateUser error", err)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	createUser.Password = hashedPassword

	user, err := u.store.CreateUser(ctx, createUser)
	if err != nil {
		u.log.Error("resources.PostV1Users", "store.CreateUser error", err)
		return nil, err
	}

	response := u.convert.PostUserAPIFromStorage(user)
	return response, nil
}

// (POST /v1/users/{id}/password)
func (u *Users) PostV1UsersIdPassword(ctx context.Context, request api.PostV1UsersIdPasswordRequestObject) (api.PostV1UsersIdPasswordResponseObject, error) {
	user, err := u.store.GetUser(ctx, request.Id)
	if err != nil {
		u.log.Error("resources.PostV1UsersIdPassword", "store.GetUser error", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request.Body.OldPassword)); err != nil {
		u.log.Error("resources.PostV1UsersIdPassword", "bcrypt.CompareHashAndPassword error", err)
		return nil, fmt.Errorf("Incorrect password")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("resources.PostV1UsersIdPassword", "bcrypt.GenerateFromPassword error", err)
		return nil, fmt.Errorf("Internal sever error")
	}

	err = u.store.UpdateUserPassword(ctx, &models.UpdateUserPassword{
		ID:          request.Id,
		NewPassword: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return api.PostV1UsersIdPassword204Response{}, nil
}
