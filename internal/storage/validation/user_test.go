package validation_test

import (
	"database/sql"
	"testing"

	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserParams(t *testing.T) {
	tests := []struct {
		name     string
		arg      db.CreateUserParams
		expected error
	}{
		{
			name: "valid user name",
			arg: db.CreateUserParams{
				Name: sql.NullString{String: "John Doe", Valid: true},
			},
			expected: nil,
		},
		{
			name: "empty user name",
			arg: db.CreateUserParams{
				Name: sql.NullString{String: "", Valid: true},
			},
			expected: validation.UserValidationError,
		},
		{
			name: "invalid user name (not valid)",
			arg: db.CreateUserParams{
				Name: sql.NullString{String: "", Valid: false},
			},
			expected: validation.UserValidationError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.UserParams(tt.arg)
			if tt.expected != nil {
				assert.Equal(t, tt.expected, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
