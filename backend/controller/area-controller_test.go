package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/controller"
	"area/schemas"
)

// MockAreaService is a mock implementation of the AreaService interface.
type MockAreaService struct {
	mock.Mock
}

func (m *MockAreaService) InitArea(area schemas.Area) {
	m.Called(area)
}

func (m *MockAreaService) CreateArea(area schemas.AreaMessage, token string) (string, error) {
	args := m.Called(area, token)
	return args.String(0), args.Error(1)
}

func (m *MockAreaService) GetUserAreas(token string) ([]schemas.Area, error) {
	args := m.Called(token)
	return args.Get(0).([]schemas.Area), args.Error(1)
}

func (m *MockAreaService) UpdateUserArea(token string, area schemas.Area) (schemas.Area, error) {
	args := m.Called(token, area)
	return args.Get(0).(schemas.Area), args.Error(1)
}

func (m *MockAreaService) DeleteUserArea(token string, area struct{ Id uint64 }) (schemas.Area, error) {
	args := m.Called(token, area)
	return args.Get(0).(schemas.Area), args.Error(1)
}

func (m *MockAreaService) AreaExist(areaID uint64) bool {
	args := m.Called(areaID)
	return args.Bool(0)
}

func (m *MockAreaService) FindAll() ([]schemas.Area, error) {
	args := m.Called()
	return args.Get(0).([]schemas.Area), args.Error(1)
}

func TestCreateArea(t *testing.T) {
	mockService := new(MockAreaService)
	controller := controller.NewAreaController(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/create", func(ctx *gin.Context) {
		result, err := controller.CreateArea(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": result})
	})

	areaMessage := schemas.AreaMessage{ /* fill with appropriate fields */ }
	areaMessageJSON, _ := json.Marshal(areaMessage)
	mockService.On("CreateArea", areaMessage, "test-token").Return("area-id", nil)

	req, _ := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(areaMessageJSON))
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetUserAreas(t *testing.T) {
	mockService := new(MockAreaService)
	controller := controller.NewAreaController(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/areas", func(ctx *gin.Context) {
		areas, err := controller.GetUserAreas(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, areas)
	})

	expectedAreas := []schemas.Area{ /* fill with appropriate fields */ }
	mockService.On("GetUserAreas", "test-token").Return(expectedAreas, nil)

	req, _ := http.NewRequest(http.MethodGet, "/areas", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateUserArea(t *testing.T) {
	mockService := new(MockAreaService)
	controller := controller.NewAreaController(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/update", func(ctx *gin.Context) {
		area, err := controller.UpdateUserArea(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, area)
	})

	area := schemas.Area{ /* fill with appropriate fields */ }
	areaJSON, _ := json.Marshal(area)
	mockService.On("UpdateUserArea", "test-token", area).Return(area, nil)

	req, _ := http.NewRequest(http.MethodPut, "/update", bytes.NewBuffer(areaJSON))
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteUserArea(t *testing.T) {
	mockService := new(MockAreaService)
	controller := controller.NewAreaController(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/delete", func(ctx *gin.Context) {
		area, err := controller.DeleteUserArea(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, area)
	})

	areaID := struct{ Id uint64 }{Id: 1}
	areaIDJSON, _ := json.Marshal(areaID)
	expectedArea := schemas.Area{ /* fill with appropriate fields */ }
	mockService.On("DeleteUserArea", "test-token", areaID).Return(expectedArea, nil)

	req, _ := http.NewRequest(http.MethodDelete, "/delete", bytes.NewBuffer(areaIDJSON))
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
