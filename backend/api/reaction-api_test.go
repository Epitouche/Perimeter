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

// MockReactionController is a mock implementation of the ReactionController interface
type MockReactionController struct {
	mock.Mock
}

func (m *MockReactionController) GetReactionsInfo(serviceID uint64) ([]schemas.Reaction, error) {
	args := m.Called(serviceID)
	return args.Get(0).([]schemas.Reaction), args.Error(1)
}

func (m *MockReactionController) GetReactionByReactionID(reactionID uint64) (schemas.Reaction, error) {
	args := m.Called(reactionID)
	return args.Get(0).(schemas.Reaction), args.Error(1)
}

func TestGetReactionsInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := new(MockReactionController)
	router := gin.Default()
	apiRoutes := router.Group("/")
	mockUserService := new(test.MockUserService)
	api.NewReactionApi(mockController, apiRoutes, mockUserService)

	t.Run("valid request", func(t *testing.T) {
		serviceID := uint64(1)
		mockReactions := []schemas.Reaction{{Id: 1, Name: "Test Reaction"}}
		mockController.On("GetReactionsInfo", serviceID).Return(mockReactions, nil)

		req, _ := http.NewRequest(http.MethodGet, "/reaction/info/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("invalid service ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/reaction/info/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		serviceID := uint64(1)
		mockController.On("GetReactionsInfo", serviceID).Return(nil, assert.AnError)

		req, _ := http.NewRequest(http.MethodGet, "/reaction/info/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
