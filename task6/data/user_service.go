package data

import (
	"context"
	"errors"
	"fmt"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func init() {
	userCollection = client.Database(dbName).Collection("users")
}

func Register(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser = userCollection.FindOne(ctx, bson.M{"username": user.Username})
	fmt.Println(existingUser)
	if existingUser.Err() == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.ID = primitive.NewObjectID()
	_, err = userCollection.InsertOne(ctx, user)
	return err
}

func Authenticate(username, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid username or password")
	}

	return user, nil
}

func GetUser(userID primitive.ObjectID) *models.User {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := taskCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return &models.User{}
	}

	return &user
}
