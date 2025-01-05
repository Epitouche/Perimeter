package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/api"
	"area/schemas"
)

type MockDropboxController struct {
	mock.Mock
}

func (m *MockDropboxController) RedirectToService(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockDropboxController) HandleServiceCallback(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockDropboxController) HandleServiceCallbackMobile(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockDropboxController) GetUserInfo(ctx *gin.Context) (schemas.UserCredentials, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.UserCredentials), args.Error(1)
}

// DROPBOX CONTROLLER MOCK.

func (m *MockDropboxController) GetUserFile(ctx *gin.Context) (userFile []schemas.DropboxFile, err error) {
	args := m.Called(ctx)
	userFile = args.Get(0).([]schemas.DropboxFile)
	err = args.Error(1)
	return userFile, err
}

func TestDropboxAPI(t *testing.T) {
	t.Parallel()

	mockController := new(MockDropboxController)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	api.NewDropboxAPI(mockController, apiRoutes)

	t.Run("TestRedirectToService", func(t *testing.T) {
		t.Parallel()
		mockController.On("RedirectToService", mock.Anything).Return("http://example.com/auth", nil)

		responseRecorder := httptest.NewRecorder()
		ctx := context.Background()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/dropbox/auth", nil)
		router.ServeHTTP(responseRecorder, req)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.Contains(t, responseRecorder.Body.String(), "http://example.com/auth")
	})
	t.Run("TestHandleServiceCallback", func(t *testing.T) {
		t.Parallel()

		mockController.On("HandleServiceCallback", mock.Anything).Return("mock_token", nil)

		responseRecorder := httptest.NewRecorder()
		ctx := context.Background()

		req, _ := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/api/dropbox/auth/callback",
			nil,
		)
		router.ServeHTTP(responseRecorder, req)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.Contains(t, responseRecorder.Body.String(), "mock_token")
	})
	t.Run("TestHandleServiceCallback", func(t *testing.T) {
		t.Parallel()

		mockController.On("HandleServiceCallbackMobile", mock.Anything).
			Return("mock_mobile_token", nil)

		responseRecorder := httptest.NewRecorder()
		ctx := context.Background()
		req, _ := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/api/dropbox/auth/callback/mobile",
			nil,
		)
		router.ServeHTTP(responseRecorder, req)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.Contains(t, responseRecorder.Body.String(), "mock_mobile_token")
	})
	// t.Run("TestHandleServiceCallback", func(t *testing.T) {
	// 	mockUserInfo := schemas.UserCredentials{Username: "Test User", Email: "aze"}
	// 	mockController.On("GetUserInfo", mock.Anything).Return(mockUserInfo, nil)

	// 	responseRecorder := httptest.NewRecorder()
	// 	req, _ := http.NewRequest(http.MethodGet, "/api/dropbox/info", nil)
	// 	router.ServeHTTP(responseRecorder, req)

	// 	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	// 	assert.Contains(t, responseRecorder.Body.String(), "Test User")
	// })
}
