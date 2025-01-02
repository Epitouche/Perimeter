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

type MockActionController struct {
	mock.Mock
}

func (m *MockActionController) GetActionsInfo(id uint64) ([]schemas.Action, error) {
	args := m.Called(id)
	return args.Get(0).([]schemas.Action), args.Error(1)
}

func TestGetActionsInfo(t *testing.T) {
	t.Parallel()

	if gin.Mode() != gin.TestMode {
		gin.SetMode(gin.TestMode)
	}

	mockController := new(MockActionController)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	api.NewActionApi(mockController, apiRoutes) // Assuming NewActionApi registers routes

	// t.Run("Success", func(t *testing.T) {
	// 	mockActions := []schemas.Action{
	// 		{Id: 1, Name: "Action1"},
	// 		{Id: 2, Name: "Action2"},
	// 	}
	// 	mockController.On("GetActionsInfo", uint64(1)).Return(mockActions, nil)

	// 	req, _ := http.NewRequest(http.MethodGet, "/api/action/info/1", nil)
	// 	w := httptest.NewRecorder()
	// 	router.ServeHTTP(w, req)

	// 	assert.Equal(t, http.StatusOK, w.Code)
	// 	mockController.AssertExpectations(t)
	// })

	// t.Run("Invalid ID", func(t *testing.T) {
	// 	req, _ := http.NewRequest(http.MethodGet, "/api/action/info/invalid", nil)
	// 	w := httptest.NewRecorder()
	// 	router.ServeHTTP(w, req)

	// 	assert.Equal(t, http.StatusBadRequest, w.Code)
	// })

	// t.Run("Internal Server Error", func(t *testing.T) {
	// 	mockController.On("GetActionsInfo", uint64(2)).Return(nil, assert.AnError)

	// 	req, _ := http.NewRequest(http.MethodGet, "/api/action/info/2", nil)
	// 	w := httptest.NewRecorder()
	// 	router.ServeHTTP(w, req)

	// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// 	mockController.AssertExpectations(t)
	// })

	t.Run("Unauthorized", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/action/info/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		mockController.AssertExpectations(t)
	})
}
