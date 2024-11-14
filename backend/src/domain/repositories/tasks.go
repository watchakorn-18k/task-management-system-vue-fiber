package repositories

import (
	"context"
	"os"
	ds "task_management_system/src/domain/datasources"
	"task_management_system/src/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepo struct {
	Collection *mongo.Collection
	Context    context.Context
}

type ITaskRepository interface {
	AddNewTask(data *entities.TaskModel) error
	GetAllTasks() (*[]entities.TaskModel, error)
	GetTask(taskID string) (*entities.TaskModel, error)
	EditTask(taskID string, data *entities.TaskModel) error
	DeleteTask(taskID string) error
}

func NewTaskRepository(db *ds.MongoDB) ITaskRepository {
	collection := db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("tasks")
	return &taskRepo{
		Collection: collection,
		Context:    db.Context,
	}
}

func (repo *taskRepo) AddNewTask(data *entities.TaskModel) error {
	_, err := repo.Collection.InsertOne(repo.Context, data)
	if err != nil {
		return err
	}
	return nil
}

func (repo *taskRepo) GetAllTasks() (*[]entities.TaskModel, error) {
	var tasks []entities.TaskModel
	cursor, err := repo.Collection.Find(repo.Context, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(repo.Context)
	for cursor.Next(repo.Context) {
		var task entities.TaskModel
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return &tasks, nil
}

func (repo *taskRepo) GetTask(taskID string) (*entities.TaskModel, error) {
	var task entities.TaskModel
	err := repo.Collection.FindOne(repo.Context, bson.M{"_id": taskID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (repo *taskRepo) EditTask(taskID string, data *entities.TaskModel) error {
	filter := bson.M{"_id": taskID}
	update := bson.M{"$set": data}
	_, err := repo.Collection.UpdateOne(repo.Context, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repo *taskRepo) DeleteTask(taskID string) error {
	filter := bson.M{"_id": taskID}
	_, err := repo.Collection.DeleteOne(repo.Context, filter)
	if err != nil {
		return err
	}
	return nil
}
