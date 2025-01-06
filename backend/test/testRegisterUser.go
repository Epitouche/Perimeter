package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func RegisterUser(t *testing.T, router *gin.Engine) (bearerToken string) {
	t.Helper() // Mark this function as a test helper

	// Define the raw JSON body for the test
	requestBody := `{
			"username": "toto",
			"email": "test@gmail.com",
			"password": "totototo"
		}`
	reqBody := bytes.NewBufferString(requestBody)

	// Perform the HTTP POST request
	responseRecorder := httptest.NewRecorder()
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/user/register", reqBody)
	require.NoError(t, err, "failed to create request")
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, responseRecorder.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	require.NoError(t, err, "failed to parse response JSON")

	// Assert the token exists and is non-empty
	token, exists := response["token"]
	assert.True(t, exists, "response does not contain 'token' key")
	assert.IsType(t, "", token, "token is not a string")
	assert.NotEmpty(t, token, "token should not be empty")
	bearerToken = "Bearer " + token.(string)

	return bearerToken
}

func LoginUser(t *testing.T, router *gin.Engine) (bearerToken string) {
	t.Helper() // Mark this function as a test helper

	// Define the raw JSON body for the test
	requestBody := `{
        "username": "toto",
        "password": "totototo"
    }`
	reqBody := bytes.NewBufferString(requestBody)

	// Create a context
	ctx := context.Background()

	// Perform the HTTP POST request
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/user/login", reqBody)
	require.NoError(t, err, "failed to create request")
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(responseRecorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	require.NoError(t, err, "failed to parse response JSON")

	// Assert the token exists and is non-empty
	token, exists := response["token"]
	assert.True(t, exists, "response does not contain 'token' key")
	assert.IsType(t, "", token, "token is not a string")
	assert.NotEmpty(t, token, "token should not be empty")

	bearerToken = "Bearer " + token.(string)

	return bearerToken
}
