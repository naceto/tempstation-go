package convert

import (
	api "github.com/naceto/tempstation/internal/generated/api/users"
	"github.com/naceto/tempstation/internal/storage/models"
)

// goverter:converter
type Convert interface {
	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Body.Name Name
	// goverter:map Body.Email Email
	CreateUserApiModelToDbModel(input api.PostV1UsersRequestObject) models.CreateUser
}
