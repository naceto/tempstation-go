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
	CreateUserAPIModelToDbModel(input *api.PostV1UsersRequestObject) *models.CreateUser

	// goverter:map . UserResponseJSONResponse
	CreateUserDbModelToAPIModel(input *models.User) *api.PostV1Users200JSONResponse

	// goverter:map ID Id
	CreateUserDbModelToAPIModelEmbedded(input models.User) api.UserResponseJSONResponse
}
