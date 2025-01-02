package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"area/schemas"
	"area/test"
)

func TestPingRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()
	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/ping", nil)
	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")
	assert.JSONEq(
		t,
		`{"message":"pong"}`,
		responseRecorder.Body.String(),
		"unexpected response body",
	)
}

func TestAboutJsonRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/about.json", nil)
	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	require.NoError(t, err, "failed to parse response JSON")

	// Assert the client exists and is non-empty
	client, exists := response["client"]
	assert.True(t, exists, "response does not contain 'client' key")
	assert.NotEmpty(t, client, "client should not be empty")

	// Assert the server exists and is non-empty
	server, exists := response["server"]
	assert.True(t, exists, "response does not contain 'server' key")
	assert.NotEmpty(t, server, "server should not be empty")
}

func TestNotFoundRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/no-route", nil)
	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code, "unexpected HTTP status code")
}

func TestBackendPortNotSet(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests

	// Set up the router (defined in main.go)
	assert.Panics(t, func() { setupRouter() }, "expected panic")
}

func TestGmailRedirectToServiceRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/gmail/auth", nil)
	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	require.NoError(t, err, "failed to parse response JSON")

	// Assert the authentication_url exists and is non-empty
	authentication_url, exists := response["authentication_url"]
	assert.True(t, exists, "response does not contain 'authentication_url' key")
	assert.IsType(t, "", authentication_url, "authentication_url is not a string")
	assert.NotEmpty(t, authentication_url, "authentication_url should not be empty")
}

func TestSpotifyRedirectToServiceRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/spotify/auth", nil)
	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	require.NoError(t, err, "failed to parse response JSON")

	// Assert the authentication_url exists and is non-empty
	authentication_url, exists := response["authentication_url"]
	assert.True(t, exists, "response does not contain 'authentication_url' key")
	assert.IsType(t, "", authentication_url, "authentication_url is not a string")
	assert.NotEmpty(t, authentication_url, "authentication_url should not be empty")
}

func TestGithubRedirectToServiceRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/github/auth", nil)
	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	require.NoError(t, err, "failed to parse response JSON")

	// Assert the authentication_url exists and is non-empty
	authentication_url, exists := response["authentication_url"]
	assert.True(t, exists, "response does not contain 'authentication_url' key")
	assert.IsType(t, "", authentication_url, "authentication_url is not a string")
	assert.NotEmpty(t, authentication_url, "authentication_url should not be empty")
}

func TestDropboxRedirectToServiceRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Set up the router (defined in main.go)
	router := setupRouter()

	// Perform the HTTP request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/dropbox/auth", nil)
	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	require.NoError(t, err, "failed to parse response JSON")

	// Assert the authentication_url exists and is non-empty
	authentication_url, exists := response["authentication_url"]
	assert.True(t, exists, "response does not contain 'authentication_url' key")
	assert.IsType(t, "", authentication_url, "authentication_url is not a string")
	assert.NotEmpty(t, authentication_url, "authentication_url should not be empty")
}

func TestRegisterUserRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	ctx := context.Background()

	// Create Postgres container
	postgresContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := postgresContainer.Terminate(ctx)
		require.NoError(t, err)
	}()

	// Set up the router (defined in main.go)
	router := setupRouter()

	test.RegisterUser(t, router)
}

func TestLoginUserRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	t.Run("no registered user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		// Define the raw JSON body for the test
		requestBody := `{
			"username": "toto",
			"password": "totototo"
		}`
		reqBody := bytes.NewBuffer([]byte(requestBody))

		// Perform the HTTP POST request
		responseRecorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/user/login", reqBody)
		require.NoError(t, err, "failed to create request")
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(responseRecorder, req)

		// Assert the response
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "unexpected HTTP status code")

		// Parse and validate the response JSON
		var response map[string]interface{}
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response JSON")

		// Assert the error exists and is non-empty
		error, exists := response["error"]
		assert.True(t, exists, "response does not contain 'error' key")
		assert.IsType(t, "", error, "error is not a string")
		assert.NotEmpty(t, error, "error should not be empty")
	})

	t.Run("registered user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		test.RegisterUser(t, router)

		test.LoginUser(t, router)
	})
}

func TestActionRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	t.Run("no registered user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		// Perform the HTTP POST request
		responseRecorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/action/info/1", nil)
		require.NoError(t, err, "failed to create request")
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(responseRecorder, req)

		// Assert the response
		assert.Equal(
			t,
			http.StatusUnauthorized,
			responseRecorder.Code,
			"unexpected HTTP status code",
		)

		// Parse and validate the response JSON
		var response map[string]interface{}
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response JSON")

		// Assert the error exists and is non-empty
		error, exists := response["error"]
		assert.True(t, exists, "response does not contain 'error' key")
		assert.IsType(t, "", error, "error is not a string")
		assert.NotEmpty(t, error, "error should not be empty")
	})

	t.Run("registered user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		bearerToken := test.RegisterUser(t, router)

		// Perform the HTTP POST request
		responseRecorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/action/info/1", nil)
		require.NoError(t, err, "failed to create request")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", bearerToken)

		router.ServeHTTP(responseRecorder, req)

		// Assert the response
		assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

		// Parse and validate the response JSON
		var response []schemas.Action
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response JSON")

		// Assert the message exists and is non-empty
		// message, exists := response["message"]
		// assert.True(t, exists, "response does not contain 'message' key")
		// assert.IsType(t, "", message, "message is not a string")
		// assert.NotEmpty(t, message, "message should not be empty")
	})

	t.Run("registered user, id not a number", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		bearerToken := test.RegisterUser(t, router)

		// Perform the HTTP POST request
		responseRecorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/action/info/test", nil)
		require.NoError(t, err, "failed to create request")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", bearerToken)

		router.ServeHTTP(responseRecorder, req)

		// Assert the response
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "unexpected HTTP status code")

		// Parse and validate the response JSON
		var response map[string]interface{}
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response JSON")

		// Assert the error exists and is non-empty
		error, exists := response["error"]
		assert.True(t, exists, "response does not contain 'error' key")
		assert.IsType(t, "", error, "error is not a string")
		assert.NotEmpty(t, error, "error should not be empty")
	})
}

func TestReactionRoute(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests
	t.Run("no registered user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		// Perform the HTTP POST request
		responseRecorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/reaction/info/1", nil)
		require.NoError(t, err, "failed to create request")
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(responseRecorder, req)

		// Assert the response
		assert.Equal(
			t,
			http.StatusUnauthorized,
			responseRecorder.Code,
			"unexpected HTTP status code",
		)

		// Parse and validate the response JSON
		var response map[string]interface{}
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response JSON")

		// Assert the error exists and is non-empty
		error, exists := response["error"]
		assert.True(t, exists, "response does not contain 'error' key")
		assert.IsType(t, "", error, "error is not a string")
		assert.NotEmpty(t, error, "error should not be empty")
	})

	t.Run("registered user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		bearerToken := test.RegisterUser(t, router)

		// Perform the HTTP POST request
		responseRecorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/reaction/info/1", nil)
		require.NoError(t, err, "failed to create request")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", bearerToken)

		router.ServeHTTP(responseRecorder, req)

		// Assert the response
		assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

		// Parse and validate the response JSON
		var response []schemas.Reaction
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response JSON")

		// Assert the message exists and is non-empty
		// message, exists := response["message"]
		// assert.True(t, exists, "response does not contain 'message' key")
		// assert.IsType(t, "", message, "message is not a string")
		// assert.NotEmpty(t, message, "message should not be empty")
	})

	t.Run("registered user, id not a number", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create Postgres container
		postgresContainer, err := test.CreatePostgresContainer(ctx)
		require.NoError(t, err, "failed to create Postgres container")
		assert.NotNil(t, postgresContainer, "failed to create Postgres container")

		// Clean up the container after the test
		defer func() {
			err := postgresContainer.Terminate(ctx)
			require.NoError(t, err)
		}()

		// Set up the router (defined in main.go)
		router := setupRouter()

		bearerToken := test.RegisterUser(t, router)

		// Perform the HTTP POST request
		responseRecorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			"/api/v1/reaction/info/test",
			nil,
		)
		require.NoError(t, err, "failed to create request")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", bearerToken)

		router.ServeHTTP(responseRecorder, req)

		// Assert the response
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "unexpected HTTP status code")

		// Parse and validate the response JSON
		var response map[string]interface{}
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response JSON")

		// Assert the error exists and is non-empty
		error, exists := response["error"]
		assert.True(t, exists, "response does not contain 'error' key")
		assert.IsType(t, "", error, "error is not a string")
		assert.NotEmpty(t, error, "error should not be empty")
	})
}
