package repository

import (
	"context"

	"GoWithMongo/datstore/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	CreateUser(user entity.User) error
	DeleteUser(username string) error
	GetUserByUsername(username string) (*entity.User, error)
}

// userRepositoryImpl implements UserRepository interface
type userRepositoryImpl struct {
	UserCollection *mongo.Collection
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepositoryImpl{
		UserCollection: collection,
	}
}

func (repo *userRepositoryImpl) CreateUser(user entity.User) error {
	_, err := repo.UserCollection.InsertOne(context.Background(), user)
	return err
}

func (repo *userRepositoryImpl) DeleteUser(username string) error {
	_, err := repo.UserCollection.DeleteOne(context.Background(), bson.M{"username": username})
	return err
}

func (repo *userRepositoryImpl) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := repo.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
