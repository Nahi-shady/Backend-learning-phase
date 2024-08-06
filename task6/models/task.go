package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" binding:"required,min=3,max=20"`
	Description string             `json:"description" bson:"description"`
	DueDate     string             `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
	UserID      string             `json:"user_id" bson:"user_id"`
}
