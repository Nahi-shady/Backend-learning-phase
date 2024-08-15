package repositories

import (
	"context"
	"task-manager/internal/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(task *entities.Task) error
	GetTasksByUserID(userID primitive.ObjectID) ([]entities.Task, error)
	UpdateTask(taskID primitive.ObjectID, title string, completed bool) error
	DeleteTask(taskID primitive.ObjectID) error
}

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) *MongoTaskRepository {
	return &MongoTaskRepository{
		collection: db.Collection("tasks"),
	}
}

func (r *MongoTaskRepository) CreateTask(task *entities.Task) error {
	_, err := r.collection.InsertOne(context.TODO(), task)
	return err
}

func (r *MongoTaskRepository) GetTasksByUserID(userID primitive.ObjectID) ([]entities.Task, error) {
	var tasks []entities.Task
	cursor, err := r.collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task entities.Task
		if err = cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *MongoTaskRepository) UpdateTask(taskID primitive.ObjectID, title string, completed bool) error {
	_, err := r.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": taskID},
		bson.M{"$set": bson.M{"title": title, "completed": completed, "updated_at": time.Now()}},
	)
	return err
}

func (r *MongoTaskRepository) DeleteTask(taskID primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": taskID})
	return err
}
