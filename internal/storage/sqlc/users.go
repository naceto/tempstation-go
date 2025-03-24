package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/models"
)

var ErrUserNotFound = errors.New("user not found")

// CreateUser creates a new user in the storage. It first validates the input parameters
// using the validation.UserParams function, and if the validation passes, it calls
// the underlying db.CreateUser function to create the user.
func (s *storage) CreateUser(ctx context.Context, createUser *models.CreateUser) (*models.User, error) {
	params := s.convert.CreateUserModelToDB(createUser)
	u, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	user := s.convert.UserModelFromDB(u)
	return user, nil
}

func (s *storage) GetUser(ctx context.Context, id int64) (*models.User, error) {
	u, err := s.db.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, fmt.Errorf(FormatStorageError, err)
	}

	user := s.convert.UserModelFromDB(u)
	return user, nil
}

func (s *storage) UpdateUser(ctx context.Context, params *models.UpdateUser) (*models.User, error) {
	u, err := s.db.UpdateUser(ctx, db.UpdateUserParams{
		Name: params.Name,
	})
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	user := s.convert.UserModelFromDB(u)
	return user, nil
}

// DeleteUser deletes the user with the specified ID from the storage.
func (s *storage) DeleteUser(ctx context.Context, id int64) error {
	err := s.db.DeleteUser(ctx, id)
	return err
}

// ListUsers returns a list of users paginated by <params>.
func (s *storage) ListUsers(ctx context.Context, params *models.ListUsersParams) ([]*models.ListUser, error) {
	var limit int32 = DefaultLimit
	if params.Limit != 0 {
		limit = params.Limit
	}

	u, err := s.db.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, fmt.Errorf(FormatStorageError, err)
	}

	users := s.convert.UserModelsFromDB(u)
	return users, nil
}
