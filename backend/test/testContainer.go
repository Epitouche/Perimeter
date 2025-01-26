package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreatePostgresContainer(t *testing.T, ctx context.Context) (testcontainers.Container, error) {
	t.Helper()

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Create a new container with PostgreSQL
	dbName := dbname
	dbUser := user
	dbPassword := password

	postgresContainer, err := postgres.Run(ctx,
		"postgres:17-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(20*time.Second)),
	)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		t.Fatalf("failed to get container host: %s", err)
	}

	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("failed to get container port: %s", err)
	}

	err = os.Setenv("DB_HOST", host)
	if err != nil {
		t.Fatalf("failed to set DB_HOST environment variable: %s", err)
	}
	err = os.Setenv("DB_PORT", mappedPort.Port())
	if err != nil {
		t.Fatalf("failed to set DB_PORT environment variable: %s", err)
	}

	return postgresContainer, nil
}
