package validation

import (
	"errors"

	"github.com/naceto/tempstation/internal/generated/db"
)

// Define the error as a variable
var UserValidationError = errors.New("user name cannot be empty")

// UserParams validates the parameters for creating a user
func UserParams(arg db.CreateUserParams) error {
	if !arg.Name.Valid || arg.Name.String == "" {
		return UserValidationError
	}
	// Add more validation as needed
	return nil
}
