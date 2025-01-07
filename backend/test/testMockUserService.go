package test

import (
	"github.com/stretchr/testify/mock"

	"area/schemas"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Login(
	newUser schemas.User,
) (jwtToken string, userID uint64, err error) {
	args := m.Called(jwtToken, userID, err)
	return args.String(0), args.Get(1).(uint64), args.Error(2)
}

func (m *MockUserService) Register(
	newUser schemas.User,
) (jwtToken string, userID uint64, err error) {
	args := m.Called(jwtToken, userID, err)
	return args.String(0), args.Get(1).(uint64), args.Error(2)
}

func (m *MockUserService) UpdateUserInfo(newUser schemas.User) (err error) {
	args := m.Called(newUser)
	return args.Error(0)
}

func (m *MockUserService) GetUserById(id uint64) (schemas.User, error) {
	args := m.Called(id)
	return args.Get(0).(schemas.User), args.Error(1)
}

func (m *MockUserService) GetUserInfo(token string) (schemas.User, error) {
	args := m.Called(token)
	return args.Get(0).(schemas.User), args.Error(1)
}
