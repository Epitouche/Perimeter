package database

import (
	"context"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setEnv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		panic(err)
	}
}

func unsetEnv(key string) {
	if err := os.Unsetenv(key); err != nil {
		panic(err)
	}
}

func TestConnectionWithContainer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpassword",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}
	defer postgresContainer.Terminate(ctx)

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		t.Fatalf("failed to get container host: %s", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("failed to get container port: %s", err)
	}

	setEnv("DB_HOST", host)
	setEnv("DB_PORT", port.Port())
	setEnv("POSTGRES_USER", "testuser")
	setEnv("POSTGRES_PASSWORD", "testpassword")
	setEnv("POSTGRES_DB", "testdb")

	defer unsetEnv("DB_HOST")
	defer unsetEnv("DB_PORT")
	defer unsetEnv("POSTGRES_USER")
	defer unsetEnv("POSTGRES_PASSWORD")
	defer unsetEnv("POSTGRES_DB")

	db := Connection()
	if db == nil {
		t.Fatal("Expected a valid database connection, got nil")
	}
}
