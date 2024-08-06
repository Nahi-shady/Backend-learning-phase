package main

import (
	"context"
	"fmt"
	"log"

	// "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this User type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongo")
	collection := client.Database("appdb").Collection("users")

	insertRresult, err := collection.InsertMany(context.TODO(), []interface{}{misty, ash, brock})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertRresult)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
