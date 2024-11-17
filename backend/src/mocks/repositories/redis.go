package repositories_test

import (
	"task_management_system/src/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockRedisRepository struct {
	mock.Mock
}

func (m *MockRedisRepository) GetTasks() (*[]entities.TaskModel, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*[]entities.TaskModel), nil
	}
	return args.Get(0).(*[]entities.TaskModel), args.Error(1)
}

func (m *MockRedisRepository) SetTasks(tasks []byte) error {
	args := m.Called(tasks)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return args.Error(0)
}
