package main

import (
	"context"
	"log"

	"task-manager/delivery/controllers"
	"task-manager/delivery/routers"
	infrastructures "task-manager/infrastructures"
	"task-manager/repositories"
	"task-manager/usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("task_manager_db")

	// Infrastructures services
	jwtService := infrastructures.NewJWTService("your-secret-key")

	// Repositories
	userRepo := repositories.NewMongoUserRepository(db)
	taskRepo := repositories.NewMongoTaskRepository(db)

	// Usecases
	userUsecase := usecases.NewUserUsecase(userRepo, &jwtService)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	// Controllers
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	r := routers.SetupRouter(userController, taskController, jwtService)

	r.Run(":8080")
}
