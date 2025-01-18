package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/api"
	"area/schemas"
	"area/test"
)

// MockMicrosoftController is a mock implementation of the MicrosoftController interface
type MockMicrosoftController struct {
	mock.Mock
}

func (m *MockMicrosoftController) RedirectToService(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockMicrosoftController) HandleServiceCallback(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockMicrosoftController) HandleServiceCallbackMobile(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockMicrosoftController) GetUserInfo(ctx *gin.Context) (schemas.UserCredentials, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.UserCredentials), args.Error(1)
}

func TestMicrosoftAPI(t *testing.T) {
	mockController := new(MockMicrosoftController)

	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockUserService := new(test.MockUserService)
	api.NewMicrosoftAPI(mockController, apiRoutes, mockUserService)

	t.Run("TestRedirectToService", func(t *testing.T) {
		mockController.On("RedirectToService", mock.Anything).Return("http://auth.url", nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/microsoft/auth", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "http://auth.url")
	})
	t.Run("TestHandleServiceCallback", func(t *testing.T) {
		mockController.On("HandleServiceCallback", mock.Anything).Return("mock_token", nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/microsoft/auth/callback", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "mock_token")
	})
	t.Run("TestHandleServiceCallbackMobile", func(t *testing.T) {
		mockController.On("HandleServiceCallbackMobile", mock.Anything).
			Return("mock_mobile_token", nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/microsoft/auth/callback/mobile", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "mock_mobile_token")
	})

	t.Run("TestGetUserInfoNoToken", func(t *testing.T) {
		mockUserInfo := schemas.UserCredentials{Username: "test_user", Email: "test@example.com"}

		mockController.On("GetUserInfo", mock.Anything).Return(mockUserInfo, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/microsoft/info/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "\"error\":\"No token provided\"")
	})
}
