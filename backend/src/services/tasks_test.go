package services_test

import (
	"errors"
	"task_management_system/src/domain/entities"
	repositories_test "task_management_system/src/mocks/repositories"
	"task_management_system/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockTaskRepoTaskSV  = new(repositories_test.MockTaskRepository)
	mockRedisRepoTaskSV = new(repositories_test.MockRedisRepository)
	taskSV              = services.NewTasksService(mockTaskRepoTaskSV, mockRedisRepoTaskSV)
)

func RefreshRepoTaskSV() {
	mockTaskRepoTaskSV = new(repositories_test.MockTaskRepository)
	mockRedisRepoTaskSV = new(repositories_test.MockRedisRepository)
	taskSV = services.NewTasksService(mockTaskRepoTaskSV, mockRedisRepoTaskSV)
}

func TestAddNewTask(t *testing.T) {
	t.Run("Add new task name is required", func(t *testing.T) {
		RefreshRepoTaskSV()
		err := taskSV.AddNewTask(&entities.TaskModel{})
		assert.Error(t, err)
	})
	t.Run("Add new task Error AddNewTask", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockTaskRepoTaskSV.On("AddNewTask", mock.Anything).Return(errors.New("error"))
		err := taskSV.AddNewTask(&entities.TaskModel{Name: "test"})
		assert.Equal(t, err.Error(), "error")
	})
	t.Run("Add new task success", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockTaskRepoTaskSV.On("AddNewTask", mock.Anything).Return(nil)
		mockTaskRepoTaskSV.On("GetAllTasks").Return(&[]entities.TaskModel{}, nil)
		mockRedisRepoTaskSV.On("SetTasks", mock.Anything).Return(nil)
		err := taskSV.AddNewTask(&entities.TaskModel{Name: "test"})
		assert.NoError(t, err)
	})
}

func TestGetAllTasks(t *testing.T) {
	t.Run("Get all tasks Error GetTasks redis", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockRedisRepoTaskSV.On("GetTasks").Return(&[]entities.TaskModel{}, errors.New("error"))
		mockTaskRepoTaskSV.On("GetAllTasks").Return(&[]entities.TaskModel{}, nil)
		_, err := taskSV.GetAllTasks()
		assert.Equal(t, nil, err)
	})
	t.Run("Get all tasks success", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockRedisRepoTaskSV.On("GetTasks").Return(&[]entities.TaskModel{}, nil)
		_, err := taskSV.GetAllTasks()
		assert.NoError(t, err)
	})
}

func TestEditTask(t *testing.T) {
	t.Run("Edit task Error EditTask", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockTaskRepoTaskSV.On("EditTask", mock.Anything, mock.Anything).Return(errors.New("error"))
		err := taskSV.EditTask("test", &entities.TaskModel{})
		assert.Equal(t, err.Error(), "error")
	})
	t.Run("Edit task success", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockTaskRepoTaskSV.On("EditTask", mock.Anything, mock.Anything).Return(nil)
		mockTaskRepoTaskSV.On("GetAllTasks").Return(&[]entities.TaskModel{}, nil)
		mockRedisRepoTaskSV.On("SetTasks", mock.Anything).Return(nil)
		err := taskSV.EditTask("test", &entities.TaskModel{})
		assert.NoError(t, err)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("Delete task Error DeleteTask", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockTaskRepoTaskSV.On("DeleteTask", mock.Anything).Return(errors.New("error"))
		err := taskSV.DeleteTask("test")
		assert.Equal(t, err.Error(), "error")
	})
	t.Run("Delete task success", func(t *testing.T) {
		RefreshRepoTaskSV()
		mockTaskRepoTaskSV.On("DeleteTask", mock.Anything).Return(nil)
		mockTaskRepoTaskSV.On("GetAllTasks").Return(&[]entities.TaskModel{}, nil)
		mockRedisRepoTaskSV.On("SetTasks", mock.Anything).Return(nil)
		err := taskSV.DeleteTask("test")
		assert.NoError(t, err)
	})
}
