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

type MockTokenController struct {
	mock.Mock
}

func (m *MockTokenController) DeleteUserToken(ctx *gin.Context) (schemas.Token, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.Token), args.Error(1)
}

func TestNewTokenApi(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := new(MockTokenController)
	mockUserService := new(test.MockUserService)

	router := gin.Default()
	apiRoutes := router.Group("/api")

	tokenApi := api.NewTokenApi(mockController, apiRoutes, mockUserService)

	assert.NotNil(t, tokenApi)
}

func TestDeleteUserToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := new(MockTokenController)
	mockUserService := new(test.MockUserService)

	router := gin.Default()
	apiRoutes := router.Group("/api")

	api.NewTokenApi(mockController, apiRoutes, mockUserService)

	t.Run("Unauthorized", func(t *testing.T) {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		mockToken := &schemas.Token{}
		mockController.On("DeleteUserToken", ctx).Return(mockToken, nil)

		req, _ := http.NewRequest("DELETE", "/api/token/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
