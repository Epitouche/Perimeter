package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"

	"area/schemas"
)

type MockController struct {
	mock.Mock
}

func (m *MockController) RedirectToService(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockController) HandleServiceCallback(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockController) HandleServiceCallbackMobile(ctx *gin.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockController) GetUserInfo(ctx *gin.Context) (schemas.UserCredentials, error) {
	args := m.Called(ctx)
	return args.Get(0).(schemas.UserCredentials), args.Error(1)
}

// DROPBOX CONTROLLER MOCK.

func (m *MockController) GetUserFile(ctx *gin.Context) (userFile []schemas.DropboxFile, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]schemas.DropboxFile), args.Error(1)
}
