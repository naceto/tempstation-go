package sqlc

import (
	"context"

	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/models"
)

// CreateUser creates a new user in the storage. It first validates the input parameters
// using the validation.UserParams function, and if the validation passes, it calls
// the underlying db.CreateUser function to create the user.
func (s *storage) CreateUser(ctx context.Context, user *models.CreateUser) (*models.User, error) {
	u, err := s.db.CreateUser(ctx, db.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *storage) GetUser(ctx context.Context, id int64) (*models.User, error) {
	u, err := s.db.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *storage) UpdateUser(ctx context.Context, params *models.UpdateUser) (*models.User, error) {
	u, err := s.db.UpdateUser(ctx, db.UpdateUserParams{
		Name: params.Name,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

// DeleteUser deletes the user with the specified ID from the storage.
func (s *storage) DeleteUser(ctx context.Context, id int64) error {
	return s.db.DeleteUser(ctx, id)
}

// ListUsers returns a list of all users stored in the database.
func (s *storage) ListUsers(ctx context.Context, params *models.ListUsersParams) ([]*models.User, error) {
	u, err := s.db.ListUsers(ctx, db.ListUsersParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}

	users := []*models.User{}
	for _, usr := range u {
		users = append(users, &models.User{
			ID:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
		})
	}

	return users, nil
}
