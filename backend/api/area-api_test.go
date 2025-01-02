package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/schemas"
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

func TestAreaAPI(t *testing.T) {
	if gin.Mode() != gin.TestMode {
		gin.SetMode(gin.TestMode)
	}
	t.Parallel()
	router := gin.Default()
	apiRoutes := router.Group("/api")

	mockController := new(MockAreaController)
	NewAreaAPI(mockController, apiRoutes)

	t.Run("TestCreateAreaNoToken", func(t *testing.T) {
		t.Parallel()

		mockController.On("CreateArea", mock.Anything).Return("Area created", nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/area/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error":"No token provided"}`, w.Body.String())
	})

	// t.Run("TestGetUserAreasNoToken", func(t *testing.T) {
	// 	t.Parallel()

	// 	mockAreas := []schemas.Area{
	// 		{Id: 1, UserId: 1, ActionId: 1, ReactionId: 1, Enable: true},
	// 		{Id: 2, UserId: 1, ActionId: 1, ReactionId: 1, Enable: true},
	// 	}
	// 	mockController.On("GetUserAreas", mock.Anything).Return(mockAreas, nil)

	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest(http.MethodGet, "/api/area", nil)
	// 	router.ServeHTTP(w, req)

	// 	assert.Equal(t, http.StatusUnauthorized, w.Code)
	// 	assert.JSONEq(t, `{"error":"No token provided"}`, w.Body.String())
	// })
}
