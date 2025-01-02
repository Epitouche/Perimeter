package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"area/database"
	"area/tools"
)

func TestHashPassword(t *testing.T) {
	t.Parallel()

	password, err := tools.GenerateCSRFToken()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	println(password)
	hashedPassword, err := database.HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.NotEqual(
		t,
		password,
		hashedPassword,
		"Expected password and hashed password to be different",
	)
	assert.NotNil(t, hashedPassword, "Expected hashed password, got nil")
	assert.Len(t, hashedPassword, 60, "Expected hashed password to be 60 characters long")
}

func TestDoPasswordsMatch(t *testing.T) {
	t.Parallel()

	password, err := tools.GenerateCSRFToken()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	println(password)
	hashedPassword, err := database.HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.True(t, database.DoPasswordsMatch(hashedPassword, password), "Expected passwords to match")

	assert.False(
		t,
		database.DoPasswordsMatch(hashedPassword, "wrongpassword"),
		"Expected passwords not to match",
	)

	assert.False(t, database.DoPasswordsMatch("wronghash", password), "Expected passwords not to match")
}
