package test

import (
	"area/schemas"

	"github.com/stretchr/testify/mock"
)

type MockActionRepository struct {
	mock.Mock
}

func (m *MockActionRepository) FindAll() ([]schemas.Action, error) {
	args := m.Called()
	return args.Get(0).([]schemas.Action), args.Error(1)
}

func (m *MockActionRepository) FindById(id uint64) (schemas.Action, error) {
	args := m.Called(id)
	return args.Get(0).(schemas.Action), args.Error(1)
}

func (m *MockActionRepository) FindByServiceId(serviceId uint64) ([]schemas.Action, error) {
	args := m.Called(serviceId)
	return args.Get(0).([]schemas.Action), args.Error(1)
}

func (m *MockActionRepository) FindByName(name string) ([]schemas.Action, error) {
	args := m.Called(name)
	return args.Get(0).([]schemas.Action), args.Error(1)
}

func (m *MockActionRepository) Save(action schemas.Action) error {
	args := m.Called(action)
	return args.Error(0)
}

func (m *MockActionRepository) FindByServiceByName(serviceId uint64, serviceName string) ([]schemas.Action, error) {
	args := m.Called(serviceId, serviceName)
	return args.Get(0).([]schemas.Action), args.Error(1)
}

func (m *MockActionRepository) Delete(action schemas.Action) error {
	args := m.Called(action)
	return args.Error(0)
}

func (m *MockActionRepository) Update(action schemas.Action) error {
	args := m.Called(action)
	return args.Error(0)
}
