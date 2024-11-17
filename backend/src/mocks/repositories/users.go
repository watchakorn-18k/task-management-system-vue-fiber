package repositories_test

import (
	"task_management_system/src/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) RegisterUsers(data *entities.User) error {
	args := m.Called(data)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return args.Error(0)
}

func (m *MockUserRepository) GetUser(userID string) (entities.User, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return entities.User{}, nil
	}
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) GetUsername(username string) (*entities.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, nil
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUsers(userID string, data *entities.User) error {
	args := m.Called(userID, data)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return args.Error(0)
}
