package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/schemas"
	"area/service"
	"area/test"
)

// MockServiceAction is a mock implementation of the ServiceAction interface.
type MockServiceAction struct {
	mock.Mock
}

func (m *MockServiceAction) GetServiceActionInfo() []schemas.Action {
	args := m.Called()
	return args.Get(0).([]schemas.Action)
}

func TestNewActionService(t *testing.T) {
	mockRepo := new(test.MockActionRepository)
	mockServiceService := new(test.MockServiceService)
	mockServiceService.On("GetServices").Return([]interface{}{})

	service := service.NewActionService(mockRepo, mockServiceService)

	assert.NotNil(t, service)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(test.MockActionRepository)
	mockServiceService := new(test.MockServiceService)
	mockServiceService.On("GetServices").Return([]interface{}{})

	expectedActions := []schemas.Action{{Name: "Test Action"}}
	mockRepo.On("FindAll").Return(expectedActions, nil)

	service := service.NewActionService(mockRepo, mockServiceService)
	actions, err := service.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedActions, actions)
}

func TestFindById(t *testing.T) {
	mockRepo := new(test.MockActionRepository)
	mockServiceService := new(test.MockServiceService)
	mockServiceService.On("GetServices").Return([]interface{}{})

	expectedAction := schemas.Action{Name: "Test Action"}
	mockRepo.On("FindById", uint64(1)).Return(expectedAction, nil)

	service := service.NewActionService(mockRepo, mockServiceService)
	action, err := service.FindById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAction, action)
}

func TestGetAllServicesByServiceId(t *testing.T) {
	mockRepo := new(test.MockActionRepository)
	mockServiceService := new(test.MockServiceService)
	mockServiceService.On("GetServices").Return([]interface{}{})

	expectedActions := []schemas.Action{{Name: "Test Action", Description: "Test Description"}}
	mockRepo.On("FindByServiceId", uint64(1)).Return(expectedActions, nil)

	service := service.NewActionService(mockRepo, mockServiceService)
	actions := service.GetAllServicesByServiceId(1)

	expectedActionJSON := []schemas.ActionJSON{{Name: "Test Action", Description: "Test Description"}}
	assert.Equal(t, expectedActionJSON, actions)
}

func TestSaveAllAction(t *testing.T) {
	mockRepo := new(test.MockActionRepository)
	mockServiceService := new(test.MockServiceService)
	mockServiceAction := new(MockServiceAction)

	expectedActions := []schemas.Action{{Name: "Test Action"}}
	mockServiceAction.On("GetServiceActionInfo").Return(expectedActions)
	mockServiceService.On("GetServices").Return([]interface{}{mockServiceAction})

	mockRepo.On("FindByName", "Test Action").Return([]schemas.Action{}, nil)
	mockRepo.On("Save", expectedActions[0]).Return(nil)

	service := service.NewActionService(mockRepo, mockServiceService)
	service.SaveAllAction()

	mockRepo.AssertCalled(t, "FindByName", "Test Action")
	mockRepo.AssertCalled(t, "Save", expectedActions[0])
}

func TestGetActionsInfo(t *testing.T) {
	mockRepo := new(test.MockActionRepository)
	mockServiceService := new(test.MockServiceService)
	mockServiceService.On("GetServices").Return([]interface{}{})

	expectedActions := []schemas.Action{{Name: "Test Action"}}
	mockRepo.On("FindByServiceId", uint64(1)).Return(expectedActions, nil)

	service := service.NewActionService(mockRepo, mockServiceService)
	actions, err := service.GetActionsInfo(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedActions, actions)
}
