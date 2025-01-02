package database

import (
	"context"
	"os"
	"testing"

	"area/test"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpassword")
	os.Setenv("POSTGRES_DB", "testdb")

	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	assert.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		assert.NoError(t, err)
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Connection() panicked: %v", r)
		}
	}()

	db := Connection()
	if db == nil {
		t.Error("Expected a valid database connection, got nil")
	}

	// Clean up environment variables after test
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("POSTGRES_DB")
}

func TestConnectionMissingEnvVars(t *testing.T) {
	// Save original environment variables
	originalEnv := map[string]string{
		"DB_HOST":           os.Getenv("DB_HOST"),
		"DB_PORT":           os.Getenv("DB_PORT"),
		"POSTGRES_USER":     os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_DB":       os.Getenv("POSTGRES_DB"),
	}

	// Clean up environment variables for testing missing cases
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("POSTGRES_DB")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to missing environment variables, but did not panic")
		}

		// Restore original environment variables
		for key, value := range originalEnv {
			os.Setenv(key, value)
		}
	}()

	Connection()
}
