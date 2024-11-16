package repositories

import (
	"context"
	"os"
	ds "task_management_system/src/domain/datasources"
	"task_management_system/src/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepo struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IUsersRepository interface {
	RegisterUsers(data *entities.User) error
	GetUser(userID string) (*entities.User, error)
	UpdateUsers(userID string, data *entities.User) error
}

func NewUsersRepository(db *ds.MongoDB) IUsersRepository {
	collection := db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("users")
	return &usersRepo{
		Collection: collection,
		Context:    db.Context,
	}
}

func (repo *usersRepo) RegisterUsers(data *entities.User) error {
	_, err := repo.Collection.InsertOne(repo.Context, data)
	if err != nil {
		return err
	}
	return nil
}

func (repo *usersRepo) GetUser(username string) (*entities.User, error) {
	var user entities.User
	err := repo.Collection.FindOne(repo.Context, bson.M{"user_id": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *usersRepo) UpdateUsers(userID string, data *entities.User) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$set": data}
	_, err := repo.Collection.UpdateOne(repo.Context, filter, update)
	if err != nil {
		return err
	}
	return nil
}
