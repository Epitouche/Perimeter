package database_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"area/database"
	"area/test"
)

func TestConnection(t *testing.T) {
	// Set environment variables for testing
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("POSTGRES_USER", "testuser")
	t.Setenv("POSTGRES_PASSWORD", "testpassword")
	t.Setenv("POSTGRES_DB", "testdb")

	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(t, ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Connection() panicked: %v", r)
		}
	}()

	db := database.Connection()
	if db == nil {
		t.Error("Expected a valid database connection, got nil")
	}
}

func TestConnectionMissingEnvVars(t *testing.T) {
	t.Parallel()

	// Save original environment variables
	originalEnv := map[string]string{
		"DB_HOST":           os.Getenv("DB_HOST"),
		"DB_PORT":           os.Getenv("DB_PORT"),
		"POSTGRES_USER":     os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_DB":       os.Getenv("POSTGRES_DB"),
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to missing environment variables, but did not panic")
		}

		// Restore original environment variables
		for key, value := range originalEnv {
			os.Setenv(key, value)
		}
	}()

	database.Connection()
}
