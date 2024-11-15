package services

import (
	"errors"
	"fmt"
	"task_management_system/src/domain/entities"
	"task_management_system/src/domain/repositories"
	"task_management_system/src/infrastructure/generals"
)

type taskService struct {
	TaskRepo repositories.ITaskRepository
}

type ITasksService interface {
	AddNewTask(data *entities.TaskModel) error
	GetAllTasks() (*[]entities.TaskModel, error)
	EditTask(taskID string, data *entities.TaskModel) error
	DeleteTask(taskID string) error
}

func NewTasksService(taskRepo repositories.ITaskRepository) ITasksService {
	return &taskService{
		TaskRepo: taskRepo,
	}
}

func (sv *taskService) AddNewTask(data *entities.TaskModel) error {
	if data.TaskID == "" {
		data.TaskID = fmt.Sprintf("task-%s", generals.RandomText(4))
	}
	if data.Name == "" {
		return errors.New("name is required")
	}
	if data.Status == "" {
		data.Status = entities.InProgress
	}
	if err := sv.TaskRepo.AddNewTask(data); err != nil {
		return err
	}
	return nil
}

func (sv *taskService) GetAllTasks() (*[]entities.TaskModel, error) {
	return sv.TaskRepo.GetAllTasks()
}

func (sv *taskService) EditTask(taskID string, data *entities.TaskModel) error {
	if err := sv.TaskRepo.EditTask(taskID, data); err != nil {
		return err
	}
	return nil
}

func (sv *taskService) DeleteTask(taskID string) error {
	if err := sv.TaskRepo.DeleteTask(taskID); err != nil {
		return err
	}
	return nil
}
