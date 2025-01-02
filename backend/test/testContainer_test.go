package test_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"area/test"
)

func TestCreatePostgresContainer(t *testing.T) {

	// Set environment variables for the test
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("POSTGRES_USER", "testuser")
	t.Setenv("POSTGRES_PASSWORD", "testpassword")
	t.Setenv("POSTGRES_DB", "testdb")

	ctx := context.Background()

	container, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, container, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := container.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Verify the environment variables are set correctly
	// assert.NotEqual(t, "5432", os.Getenv("DB_PORT"))
}
