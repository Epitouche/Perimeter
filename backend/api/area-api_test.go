package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Epitouche/Perimeter/api"
	"github.com/Epitouche/Perimeter/schemas"
	"github.com/Epitouche/Perimeter/test"
)

type MockAreaController struct {
	mock.Mock
}

func (m *MockAreaController) CreateArea(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockAreaController) GetUserAreas(ctx *gin.Context) ([]schemas.Area, error) {
	args := m.Called(ctx)
	return args.Get(0).([]schemas.Area), args.Error(1)
}

func (m *MockAreaController) UpdateUserArea(ctx *gin.Context) (newArea schemas.Area, err error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.Area), args.Error(1)
}

func (m *MockAreaController) DeleteUserArea(ctx *gin.Context) (newArea schemas.Area, err error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.Area), args.Error(1)
}

func TestAreaAPI(t *testing.T) {
	t.Parallel()

	router := gin.Default()
	apiRoutes := router.Group("/api")

	mockController := new(MockAreaController)
	mockUserService := new(test.MockUserService)
	api.NewAreaAPI(mockController, apiRoutes, mockUserService)

	t.Run("TestCreateAreaNoToken", func(t *testing.T) {
		t.Parallel()

		mockController.On("CreateArea", mock.Anything).Return("Area created", nil)

		responseRecorder := httptest.NewRecorder()
		ctx := context.Background()

		req, _ := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/api/github.com/Epitouche/Perimeter/",
			nil,
		)
		router.ServeHTTP(responseRecorder, req)

		assert.Equal(t, http.StatusUnauthorized, responseRecorder.Code)
		assert.JSONEq(t, `{"error":"No token provided"}`, responseRecorder.Body.String())
	})

	// t.Run("TestGetUserAreasNoToken", func(t *testing.T) {
	// 	t.Parallel()

	// 	mockAreas := []schemas.Area{
	// 		{Id: 1, UserId: 1, ActionId: 1, ReactionId: 1, Enable: true},
	// 		{Id: 2, UserId: 1, ActionId: 1, ReactionId: 1, Enable: true},
	// 	}
	// 	mockController.On("GetUserAreas", mock.Anything).Return(mockAreas, nil)

	// 	responseRecorder := httptest.NewRecorder()
	// 	req, _ := http.NewRequest(http.MethodGet, "/api/area", nil)
	// 	router.ServeHTTP(responseRecorder, req)

	// 	assert.Equal(t, http.StatusUnauthorized, responseRecorder.Code)
	// 	assert.JSONEq(t, `{"error":"No token provided"}`, responseRecorder.Body.String())
	// })
}
