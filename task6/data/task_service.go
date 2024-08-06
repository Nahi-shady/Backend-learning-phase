package data

import (
	"context"
	"log"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"

const (
	dbName         = "taskdb"
	collectionName = "tasks"
)

var (
	client         *mongo.Client
	taskCollection *mongo.Collection
)

func init() {
	var err error
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	taskCollection = client.Database(dbName).Collection(collectionName)

}

func AddTask(task *models.Task) (*mongo.InsertOneResult, error) {
	task.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := taskCollection.InsertOne(ctx, task)
	return result, err
}

func GetTaskByIDAndUserID(id, userID string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = taskCollection.FindOne(ctx, bson.M{"_id": objID, "user_id": userID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTaskByIDAndUserID(id, userID string, updatedTask *models.Task) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	result, err := taskCollection.UpdateOne(ctx, bson.M{"_id": objID, "user_id": userID}, update)
	return result, err
}

func DeleteTaskByIDAndUserID(id, userID string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := taskCollection.DeleteOne(ctx, bson.M{"_id": objID, "user_id": userID})
	return result, err
}

func GetAllTasksByUserID(userID string) ([]*models.Task, error) {
	var tasks []*models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user *models.User
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	user = GetUser(objID)
	var cursor *mongo.Cursor
	if user.Role == "Admin" {
		cursor, err = taskCollection.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
	} else {
		cursor, err = taskCollection.Find(ctx, bson.M{"_id": userID})
		if err != nil {
			return nil, err
		}
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
