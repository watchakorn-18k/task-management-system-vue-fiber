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
}

func NewTasksService(taskRepo repositories.ITaskRepository) ITasksService {
	return &taskService{
		TaskRepo: taskRepo,
	}
}

func (sv *taskService) AddNewTask(data *entities.TaskModel) error {
	data.TaskID = fmt.Sprintf("task-%s", generals.RandomText(4))
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
