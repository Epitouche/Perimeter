package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/api"
	"area/middlewares"
	"area/schemas"
	"area/test"
)

type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) Login(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockUserController) Register(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockUserController) GetUser(ctx *gin.Context) (schemas.UserCredentials, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.UserCredentials), args.Error(1)
}

func (m *MockUserController) GetUserAllInfo(ctx *gin.Context) (schemas.UserAllInfo, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.UserAllInfo), args.Error(1)
}

func (m *MockUserController) UpdateUser(ctx *gin.Context) (schemas.User, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.User), args.Error(1)
}

func (m *MockUserController) DeleteUser(ctx *gin.Context) (schemas.User, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.User), args.Error(1)
}

func TestNewUserApi(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockController := new(MockUserController)
	mockService := new(test.MockUserService)

	userApi := api.NewUserApi(mockController, apiRoutes, mockService)

	assert.NotNil(t, userApi)
}

func TestUserApi_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockController := new(MockUserController)
	mockService := new(test.MockUserService)

	userApi := api.NewUserApi(mockController, apiRoutes, mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/api/user/login", nil)

	mockController.On("Login", ctx).Return("mockToken", nil)

	userApi.Login(apiRoutes)

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserApi_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockController := new(MockUserController)
	mockService := new(test.MockUserService)

	api.NewUserApi(mockController, apiRoutes, mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/api/user/register", nil)

	mockController.On("Register", ctx).Return("mockToken", nil)

	// userApi.Register(apiRoutes)

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserApi_GetUserInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockController := new(MockUserController)
	mockService := new(test.MockUserService)

	api.NewUserApi(mockController, apiRoutes, mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/api/user/info/", nil)

	mockController.On("GetUser", ctx).Return(&schemas.UserCredentials{}, nil)

	// userApi.GetUserInfo(apiRoutes.Group("/info/", middlewares.AuthorizeJWT(mockService)))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestUserApi_GetUserAllInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockController := new(MockUserController)
	mockService := new(test.MockUserService)

	userApi := api.NewUserApi(mockController, apiRoutes, mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/api/user/info/all", nil)

	mockController.On("GetUserAllInfo", ctx).Return(&schemas.UserAllInfo{}, nil)

	userApi.GetUserAllInfo(apiRoutes.Group("/info", middlewares.AuthorizeJWT(mockService)))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestUserApi_UpdateUserInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockController := new(MockUserController)
	mockService := new(test.MockUserService)

	userApi := api.NewUserApi(mockController, apiRoutes, mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPut, "/api/user/info/", nil)

	mockController.On("UpdateUser", ctx).Return(&schemas.User{}, nil)

	userApi.UpdateUserInfo(apiRoutes.Group("/info/", middlewares.AuthorizeJWT(mockService)))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestUserApi_DeleteUserInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes := router.Group("/api")
	mockController := new(MockUserController)
	mockService := new(test.MockUserService)

	userApi := api.NewUserApi(mockController, apiRoutes, mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodDelete, "/api/user/info/", nil)

	mockController.On("DeleteUser", ctx).Return(&schemas.User{}, nil)

	userApi.DeleteUserInfo(apiRoutes.Group("/info", middlewares.AuthorizeJWT(mockService)))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
