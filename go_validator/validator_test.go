package go_validator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidator_profile(t *testing.T) {
	validate := validator.New()

	invalidProfile := Profile{
		Username: "user@name", // Contains non-alphanumeric characters
		Password: "short",     // Password is too short
	}

	if err := validate.Struct(validProfile); err != nil {
		t.Error("validation error:", err)
	}
}
