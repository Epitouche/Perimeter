package controller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"area/controller"
	"area/schemas"
)

// MockActionService is a mock implementation of the ActionService interface for testing purposes.
type MockActionService struct {
	mock.Mock
}

func (m *MockActionService) GetActionsInfo(id uint64) ([]schemas.Action, error) {
	args := m.Called(id)
	return args.Get(0).([]schemas.Action), args.Error(1)
}

func (m *MockActionService) SaveAllAction() {
	m.Called()
}

func (m *MockActionService) GetAllServicesByServiceId(id uint64) []schemas.ActionJSON {
	args := m.Called(id)
	return args.Get(0).([]schemas.ActionJSON)
}

func (m *MockActionService) FindAll() ([]schemas.Action, error) {
	args := m.Called()
	return args.Get(0).([]schemas.Action), args.Error(1)
}

func (m *MockActionService) FindById(id uint64) (schemas.Action, error) {
	args := m.Called(id)
	return args.Get(0).(schemas.Action), args.Error(1)
}

func TestGetActionsInfo(t *testing.T) {
	mockService := new(MockActionService)
	controller := controller.NewActionController(mockService)

	t.Run("success", func(t *testing.T) {
		expectedActions := []schemas.Action{{Id: 1}, {Id: 2}}
		mockService.On("GetActionsInfo", uint64(1)).Return(expectedActions, nil)

		actions, err := controller.GetActionsInfo(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedActions, actions)
		mockService.AssertExpectations(t)
	})

	// t.Run("error", func(t *testing.T) {
	// 	mockService.On("GetActionsInfo", uint64(1)).Return(nil, errors.New("service error"))

	// 	actions, err := controller.GetActionsInfo(1)

	// 	assert.Error(t, err)
	// 	assert.Nil(t, actions)
	// 	mockService.AssertExpectations(t)
	// })
}

func TestGetActionByActionID(t *testing.T) {
	mockService := new(MockActionService)
	controller := controller.NewActionController(mockService)

	t.Run("success", func(t *testing.T) {
		expectedAction := schemas.Action{Id: 1}
		mockService.On("FindById", uint64(1)).Return(expectedAction, nil)

		action, err := controller.GetActionByActionID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedAction, action)
		mockService.AssertExpectations(t)
	})

	// t.Run("error", func(t *testing.T) {
	// 	mockService.On("FindById", uint64(1)).Return(schemas.Action{}, errors.New("service error"))

	// 	action, err := controller.GetActionByActionID(2)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, schemas.Action{}, action)
	// 	mockService.AssertExpectations(t)
	// })
}
