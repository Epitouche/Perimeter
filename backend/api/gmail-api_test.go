package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/test"
)

func TestGmailAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := new(test.MockController)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	NewGmailAPI(mockController, apiRoutes)

	t.Run("TestRedirectToService", func(t *testing.T) {
		mockController.On("RedirectToService", mock.Anything).Return("http://example.com/auth", nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/gmail/auth", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "http://example.com/auth")
	})
	t.Run("TestHandleServiceCallback", func(t *testing.T) {
		mockController.On("HandleServiceCallback", mock.Anything).Return("mock_token", nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/gmail/auth/callback", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "mock_token")
	})
	t.Run("TestHandleServiceCallback", func(t *testing.T) {
		mockController.On("HandleServiceCallbackMobile", mock.Anything).Return("mock_mobile_token", nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/gmail/auth/callback/mobile", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "mock_mobile_token")
	})
	// t.Run("TestHandleServiceCallback", func(t *testing.T) {
	// 	mockUserInfo := schemas.UserCredentials{Username: "Test User", Email: "aze"}
	// 	mockController.On("GetUserInfo", mock.Anything).Return(mockUserInfo, nil)

	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest("GET", "/api/gmail/info", nil)
	// 	router.ServeHTTP(w, req)

	// 	assert.Equal(t, http.StatusOK, w.Code)
	// 	assert.Contains(t, w.Body.String(), "Test User")
	// })
}
