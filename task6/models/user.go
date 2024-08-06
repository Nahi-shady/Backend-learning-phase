package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" binding:"required,min=3,max=20"`
	Password string             `json:"password" bson:"password" binding:"required,min=6,max=20"`
	Role     string             `json:"role" bson:"role"`
}
