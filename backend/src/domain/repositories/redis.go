package repositories

import (
	"context"
	"errors"
	"task_management_system/src/domain/datasources"
	"task_management_system/src/domain/entities"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

type redisRepo struct {
	Context   context.Context
	RedisWR   *redis.Client
	RedisRead *redis.Client
}

type IRedisRepository interface {
	GetTasks() (*[]entities.TaskModel, error)
	SetTasks(tasks []byte) error
}

func NewRedisRepository(redis *datasources.RedisConnection) IRedisRepository {
	return &redisRepo{
		Context:   redis.Context,
		RedisWR:   redis.RedisWR,
		RedisRead: redis.RedisRead,
	}
}

func (r *redisRepo) GetTasks() (*[]entities.TaskModel, error) {
	dataSpeakerRedis, err := r.RedisRead.Get(r.Context, "GoTasks").Result()
	if err != nil {
		log.Info("error GetTasks ", err.Error())
		return nil, errors.New("error GetTasks")
	}
	var data *[]entities.TaskModel
	if err := json.Unmarshal([]byte(dataSpeakerRedis), &data); err != nil {
		log.Info("error GetTasks ", err.Error())
		return nil, errors.New("error GetTasks: " + err.Error())
	}
	log.Info("Get users data to redis success!")
	return data, nil
}

func (r *redisRepo) SetTasks(tasks []byte) error {
	if err := r.RedisWR.Set(r.Context, "GoTasks", tasks, 0).Err(); err != nil {
		log.Info("error SetTasks ", err.Error())
		return err
	}
	log.Info("Set tasks data to redis success!")
	return nil
}
