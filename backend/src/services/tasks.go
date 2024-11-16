package services

import (
	"errors"
	"fmt"
	"task_management_system/src/domain/entities"
	"task_management_system/src/domain/repositories"
	"task_management_system/src/infrastructure/generals"

	"github.com/goccy/go-json"
)

type taskService struct {
	TaskRepo  repositories.ITaskRepository
	RedisRepo repositories.IRedisRepository
}

type ITasksService interface {
	AddNewTask(data *entities.TaskModel) error
	GetAllTasks() (*[]entities.TaskModel, error)
	EditTask(taskID string, data *entities.TaskModel) error
	DeleteTask(taskID string) error
}

func NewTasksService(taskRepo repositories.ITaskRepository, redisRepo repositories.IRedisRepository) ITasksService {
	return &taskService{
		TaskRepo:  taskRepo,
		RedisRepo: redisRepo,
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
	if err := sv.SetTasks(); err != nil {
		return err
	}
	return nil
}

func (sv *taskService) GetAllTasks() (*[]entities.TaskModel, error) {
	dataRedis, err := sv.RedisRepo.GetTasks()
	if err != nil {
		return sv.TaskRepo.GetAllTasks()
	}
	return dataRedis, nil
}

func (sv *taskService) EditTask(taskID string, data *entities.TaskModel) error {
	if err := sv.TaskRepo.EditTask(taskID, data); err != nil {
		return err
	}
	if err := sv.SetTasks(); err != nil {
		return err
	}
	return nil
}

func (sv *taskService) DeleteTask(taskID string) error {
	if err := sv.TaskRepo.DeleteTask(taskID); err != nil {
		return err
	}
	if err := sv.SetTasks(); err != nil {
		return err
	}
	return nil
}

func (sv *taskService) SetTasks() error {
	tasks, err := sv.TaskRepo.GetAllTasks()
	if err != nil {
		return err
	}
	tasksByte, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	if err := sv.RedisRepo.SetTasks(tasksByte); err != nil {
		return err
	}
	return nil
}
