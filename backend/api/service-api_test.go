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
)

type MockServiceController struct {
	mock.Mock
}

func (m *MockServiceController) AboutJSON(ctx *gin.Context) (schemas.AboutJSON, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.AboutJSON), args.Error(1)
}

func (m *MockServiceController) GetServicesInfo() ([]schemas.Service, error) {
	args := m.Called()
	return args.Get(0).([]schemas.Service), args.Error(1)
}

func (m *MockServiceController) GetServiceInfoById(id uint64) (schemas.Service, error) {
	args := m.Called(id)
	return args.Get(0).(schemas.Service), args.Error(1)
}

func TestNewServiceApi(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")

	mockController := new(MockServiceController)
	serviceApi := api.NewServiceApi(mockController, apiRoutes)

	assert.NotNil(t, serviceApi)
}

func TestGetServicesInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")

	mockController := new(MockServiceController)
	api.NewServiceApi(mockController, apiRoutes)

	mockController.On("GetServicesInfo").Return([]schemas.Service{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/service/info/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestGetServiceInfoById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")

	mockController := new(MockServiceController)
	api.NewServiceApi(mockController, apiRoutes)

	mockService := schemas.Service{}
	mockController.On("GetServiceInfoById", uint64(1)).Return(mockService, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/service/info/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}
