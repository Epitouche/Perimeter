package test

import (
	"encoding/json"

	"area/schemas"
	"area/service"

	"github.com/stretchr/testify/mock"
)

type MockServiceService struct {
	mock.Mock
}

func (m *MockServiceService) HandleServiceCallback(code string,
	authorization string,
	serviceName schemas.ServiceName,
	authGetServiceAccessToken func(code string) (schemas.Token, error),
	serviceUser service.UserService,
	getUserInfo func(token string) (userInfo schemas.User, err error),
	tokenService service.TokenService,
) (string, error) {
	args := m.Called(
		authorization,
		serviceName,
		authGetServiceAccessToken,
		serviceUser,
		getUserInfo,
		tokenService,
	)
	return args.String(0), args.Error(1)
}

func (m *MockServiceService) GetAllServices() ([]schemas.ServiceJSON, error) {
	args := m.Called()
	return args.Get(0).([]schemas.ServiceJSON), args.Error(1)
}

func (m *MockServiceService) HandleServiceCallbackMobile(code string,
	serviceName schemas.ServiceName,
	mobileTokenRequest schemas.MobileTokenRequest,
	serviceUser service.UserService,
	getUserInfo func(token string) (userInfo schemas.User, err error),
	tokenService service.TokenService,
) (string, error) {
	args := m.Called(serviceName, mobileTokenRequest, serviceUser, getUserInfo, tokenService)
	return args.String(0), args.Error(1)
}

func (m *MockServiceService) GetServices() []interface{} {
	args := m.Called()
	return args.Get(0).([]interface{})
}

func (m *MockServiceService) FindActionByName(
	name string,
) func(chan string, json.RawMessage, schemas.Area) {
	args := m.Called(name)
	return args.Get(0).(func(chan string, json.RawMessage, schemas.Area))
}

func (m *MockServiceService) FindByName(name schemas.ServiceName) schemas.Service {
	args := m.Called(name)
	return args.Get(0).(schemas.Service)
}

func (m *MockServiceService) FindAll() []schemas.Service {
	args := m.Called()
	return args.Get(0).([]schemas.Service)
}

func (m *MockServiceService) FindReactionByName(
	name string,
) func(json.RawMessage, schemas.Area) string {
	args := m.Called(name)
	return args.Get(0).(func(json.RawMessage, schemas.Area) string)
}

func (m *MockServiceService) FindServiceByName(name string) schemas.Service {
	args := m.Called(name)
	return args.Get(0).(schemas.Service)
}

func (m *MockServiceService) GetServiceById(id uint64) schemas.Service {
	args := m.Called(id)
	return args.Get(0).(schemas.Service)
}

func (m *MockServiceService) GetServicesInfo() ([]schemas.Service, error) {
	args := m.Called()
	return args.Get(0).([]schemas.Service), args.Error(1)
}

func (m *MockServiceService) RedirectToServiceOauthPage(
	serviceName schemas.ServiceName,
	state string,
	redirectUri string,
) (string, error) {
	args := m.Called(serviceName, state, redirectUri)
	return args.String(0), args.Error(1)
}
