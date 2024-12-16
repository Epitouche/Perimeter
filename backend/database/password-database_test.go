package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"area/tools"
)

func TestHashPassword(t *testing.T) {
	password, err := tools.GenerateCSRFToken()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	println(password)
	hashedPassword, err := HashPassword(password)

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
	assert.Equal(t, 60, len(hashedPassword), "Expected hashed password to be 60 characters long")
}

func TestDoPasswordsMatch(t *testing.T) {
	password, err := tools.GenerateCSRFToken()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	println(password)
	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.True(t, DoPasswordsMatch(hashedPassword, password), "Expected passwords to match")

	assert.False(
		t,
		DoPasswordsMatch(hashedPassword, "wrongpassword"),
		"Expected passwords not to match",
	)

	assert.False(t, DoPasswordsMatch("wronghash", password), "Expected passwords not to match")

	assert.False(
		t,
		DoPasswordsMatch("wronghash", "wrongpassword"),
		"Expected passwords not to match",
	)
}
