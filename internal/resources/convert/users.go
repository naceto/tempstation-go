package convert

import (
	api "github.com/naceto/tempstation/internal/generated/api/users"
	"github.com/naceto/tempstation/internal/storage/models"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:extend ListUsersAPIFromStorageExtend
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

	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Params.Offset Offset
	// goverter:map Params.Limit Limit
	ListUsersAPIToStorage(input *api.GetV1UsersRequestObject) *models.ListUsersParams

	ListUsersAPIFromStorage(input []*models.User) api.GetV1Users200JSONResponse
}

func ListUsersAPIFromStorageExtend(input []*models.User) api.GetV1Users200JSONResponse {
	users := make([]api.User, 0, len(input))
	for _, u := range input {
		users = append(users, api.User{
			Id:    &u.ID,
			Name:  &u.Name,
			Email: &u.Email,
		})
	}

	return api.GetV1Users200JSONResponse{
		UsersResponseJSONResponse: api.UsersResponseJSONResponse{
			Users: users,
		},
	}
}
