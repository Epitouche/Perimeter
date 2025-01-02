package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RegisterUser(router *gin.Engine, t *testing.T) (bearerToken string) {
	// Define the raw JSON body for the test
	requestBody := `{
			"username": "toto",
			"email": "test@gmail.com",
			"password": "totototo"
		}`
	reqBody := bytes.NewBuffer([]byte(requestBody))

	// Perform the HTTP POST request
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/user/register", reqBody)
	assert.NoError(t, err, "failed to create request")
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "failed to parse response JSON")

	// Assert the token exists and is non-empty
	token, exists := response["token"]
	assert.True(t, exists, "response does not contain 'token' key")
	assert.IsType(t, "", token, "token is not a string")
	assert.NotEmpty(t, token, "token should not be empty")
	bearerToken = "Bearer " + token.(string)

	return bearerToken
}

func LoginUser(router *gin.Engine, t *testing.T) (bearerToken string) {
	// Define the raw JSON body for the test
	requestBody := `{
		"username": "toto",
		"password": "totototo"
	}`
	reqBody := bytes.NewBuffer([]byte(requestBody))

	// Perform the HTTP POST request
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/user/login", reqBody)
	assert.NoError(t, err, "failed to create request")
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code, "unexpected HTTP status code")

	// Parse and validate the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "failed to parse response JSON")

	// Assert the token exists and is non-empty
	token, exists := response["token"]
	assert.True(t, exists, "response does not contain 'token' key")
	assert.IsType(t, "", token, "token is not a string")
	assert.NotEmpty(t, token, "token should not be empty")

	bearerToken = "Bearer " + token.(string)

	return bearerToken
}
