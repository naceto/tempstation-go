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
	CreateUserAPIToStorage(input *api.PostV1UsersRequestObject) *models.CreateUser

	// goverter:map . UserResponseJSONResponse
	PostUserAPIFromStorage(input *models.User) *api.PostV1Users200JSONResponse

	// goverter:map . UserResponseJSONResponse
	GetUserAPIFromStorage(input *models.User) *api.GetV1UsersId200JSONResponse

	// goverter:map ID Id
	UserAPIFromStorageEmbedded(input models.User) api.UserResponseJSONResponse
}
