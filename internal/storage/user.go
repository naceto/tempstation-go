package storage

import (
	"context"

	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/validation"
)

// CreateUser creates a new user in the storage. It first validates the input parameters
// using the validation.UserParams function, and if the validation passes, it calls
// the underlying db.CreateUser function to create the user.
func (s *storage) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	if err := validation.UserParams(arg); err != nil { // Use the validation function
		return db.User{}, err
	}
	return s.db.CreateUser(ctx, arg)
}

func (s *storage) GetUser(ctx context.Context, id int64) (db.User, error) {
	return s.db.GetUser(ctx, id)
}

func (s *storage) UpdateUser(ctx context.Context, arg db.UpdateUserParams) error {
	return s.db.UpdateUser(ctx, arg)
}

// DeleteUser deletes the user with the specified ID from the storage.
func (s *storage) DeleteUser(ctx context.Context, id int64) error {
	return s.db.DeleteUser(ctx, id)
}

// ListUsers returns a list of all users stored in the database.
func (s *storage) ListUsers(ctx context.Context) ([]db.User, error) {
	return s.db.ListUsers(ctx)
}
