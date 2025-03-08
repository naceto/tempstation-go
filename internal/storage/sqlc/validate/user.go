package validate

import (
	"errors"

	"github.com/naceto/tempstation/internal/storage/models"
)

var (
	ErrUserValidationError  = errors.New("user name cannot be empty")
	ErrEmailValidationError = errors.New("user email cannot be empty")
)

func CreateUser(params *models.CreateUser) error {
	if params.Name == "" {
		return ErrUserValidationError
	}

	if params.Email == "" {
		return ErrEmailValidationError
	}

	return nil
}
