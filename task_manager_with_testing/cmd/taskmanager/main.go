package main

import (
	"context"
	"log"
	"task-manager/config"
	"task-manager/internal/application/services"
	"task-manager/internal/delivery/controllers"
	"task-manager/internal/delivery/routes"
	"task-manager/internal/domain/repositories"
	"task-manager/internal/infrastructure/auth"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.LoadConfig()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.DBUri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("task_manager_db")

	jwtService := auth.NewJWTService(cfg.JWTSecret)

	userRepo := repositories.NewMongoUserRepository(db)
	taskRepo := repositories.NewMongoTaskRepository(db)

	userService := *services.NewUserService(userRepo, jwtService)
	taskService := *services.NewTaskService(taskRepo)

	userController := controllers.NewUserController(userService)
	taskController := controllers.NewTaskController(taskService)

	r := routes.SetupRouter(userController, taskController, jwtService)

	r.Run(":" + cfg.ServerPort)
}
