package main

import (
	"context"
	"log"
	"task-manager/internal/application/services"
	"task-manager/internal/delivery/controllers"
	"task-manager/internal/delivery/routers"
	"task-manager/internal/domain/repositories"
	"task-manager/internal/infrastructure/auth"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("task_manager_db")

	jwtService := auth.NewJWTService("your-secret-key")
	userRepo := repositories.NewMongoUserRepository(db)
	taskRepo := repositories.NewMongoTaskRepository(db)

	userService := services.NewUserService(userRepo, jwtService)
	taskService := services.NewTaskService(taskRepo)

	userController := controllers.NewUserController(userService)
	taskController := controllers.NewTaskController(taskService)
	router := routers.SetupRouter(userController, taskController, jwtService)

	router.Run(":8080")
}
