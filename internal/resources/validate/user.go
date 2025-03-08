package validate

import (
	"errors"

	"github.com/naceto/tempstation/internal/storage/models"
)

var (
	ErrUserValidationError  = errors.New("user name cannot be empty")
	ErrEmailValidationError = errors.New("user email cannot be empty")
)

func CreateUser(user *models.CreateUser) error {
	if user.Name == "" {
		return ErrUserValidationError
	}

	if user.Email == "" {
		return ErrEmailValidationError
	}

	return nil
}
