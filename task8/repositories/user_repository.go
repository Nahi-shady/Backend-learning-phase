package repositories

import (
	"context"
	"errors"
	"task-manager/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByUsername(username string) (*domain.User, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepository) CreateUser(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser = r.collection.FindOne(ctx, bson.M{"username": user.Username})
	if existingUser.Err() == nil {
		return errors.New("username already exists")
	}

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
