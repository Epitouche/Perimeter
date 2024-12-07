package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"area/schemas"
	"area/test"
)

func TestPingRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
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
	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code, "unexpected HTTP status code")
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String(), "unexpected response body")
}

func TestAboutJsonRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
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

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/about.json", nil)
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code, "unexpected HTTP status code")

	// TODO - Add more assertions
}

func TestGmailRedirectToServiceRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
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

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/gmail/auth", nil)
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code, "unexpected HTTP status code")
	expectedResponse := schemas.JWT{}
	err = json.NewDecoder(w.Body).Decode(&expectedResponse)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	assert.NotNil(t, expectedResponse, "unexpected response body")
}

func TestSpotifyRedirectToServiceRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
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

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/spotify/auth", nil)
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code, "unexpected HTTP status code")
	expectedResponse := schemas.JWT{}
	err = json.NewDecoder(w.Body).Decode(&expectedResponse)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	assert.NotNil(t, expectedResponse, "unexpected response body")
}
