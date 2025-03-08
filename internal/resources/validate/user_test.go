package validate_test

import (
	"testing"

	"github.com/naceto/tempstation/internal/resources/validate"
	"github.com/naceto/tempstation/internal/storage/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserParams(t *testing.T) {
	tests := []struct {
		name     string
		params   *models.CreateUser
		expected error
	}{
		{
			name: "ok",
			params: &models.CreateUser{
				Name:  "John Doe",
				Email: "johndoe@example.com",
			},
			expected: nil,
		},
		{
			name: "empty user name",
			params: &models.CreateUser{
				Name: "",
			},
			expected: validate.ErrUserValidationError,
		},
		{
			name: "empty email",
			params: &models.CreateUser{
				Name: "John Doe",
			},
			expected: validate.ErrEmailValidationError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.CreateUser(tt.params)
			if tt.expected != nil {
				assert.Equal(t, tt.expected, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
