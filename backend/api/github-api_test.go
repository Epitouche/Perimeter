package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/test"
)

func TestGithubAPI(t *testing.T) {
	// set the router to test mode
	if gin.Mode() != gin.TestMode {
		gin.SetMode(gin.TestMode)
	}

	// all tests should be run in parallel
	t.Parallel()

	mockController := new(test.MockController)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	NewGithubAPI(mockController, apiRoutes)

	t.Run("TestRedirectToService", func(t *testing.T) {
		t.Parallel()

		mockController.On("RedirectToService", mock.Anything).Return("http://example.com/auth", nil)

		w := httptest.NewRecorder()
		ctx := context.Background()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/github/auth", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "http://example.com/auth")
	})
	t.Run("TestHandleServiceCallback", func(t *testing.T) {
		t.Parallel()

		mockController.On("HandleServiceCallback", mock.Anything).Return("mock_token", nil)

		w := httptest.NewRecorder()
		ctx := context.Background()

		req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/api/github/auth/callback", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "mock_token")
	})
	t.Run("TestHandleServiceCallback", func(t *testing.T) {
		t.Parallel()

		mockController.On("HandleServiceCallbackMobile", mock.Anything).
			Return("mock_mobile_token", nil)

		w := httptest.NewRecorder()
		ctx := context.Background()

		req, _ := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/api/github/auth/callback/mobile",
			nil,
		)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "mock_mobile_token")
	})
	// t.Run("TestHandleServiceCallback", func(t *testing.T) {
	// 	mockUserInfo := schemas.UserCredentials{Username: "Test User", Email: "aze"}
	// 	mockController.On("GetUserInfo", mock.Anything).Return(mockUserInfo, nil)

	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest(http.MethodGet, "/api/github/info", nil)
	// 	router.ServeHTTP(w, req)

	// 	assert.Equal(t, http.StatusOK, w.Code)
	// 	assert.Contains(t, w.Body.String(), "Test User")
	// })
}
