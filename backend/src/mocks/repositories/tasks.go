package repositories_test

import (
	"task_management_system/src/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) AddNewTask(data *entities.TaskModel) error {
	args := m.Called(data)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllTasks() (*[]entities.TaskModel, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*[]entities.TaskModel), nil
	}
	return args.Get(0).(*[]entities.TaskModel), args.Error(1)
}

func (m *MockTaskRepository) GetTask(taskID string) (*entities.TaskModel, error) {
	args := m.Called(taskID)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.TaskModel), nil
	}
	return args.Get(0).(*entities.TaskModel), args.Error(1)
}

func (m *MockTaskRepository) EditTask(taskID string, data *entities.TaskModel) error {
	args := m.Called(taskID, data)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(taskID string) error {
	args := m.Called(taskID)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return args.Error(0)
}
