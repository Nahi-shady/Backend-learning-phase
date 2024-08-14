package routers

import (
	"task-manager/internal/delivery/controllers"
	"task-manager/internal/infrastructure/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, taskController *controllers.TaskController, jwtService auth.JWTService) *gin.Engine {
	r := gin.Default()

	authMiddleware := auth.AuthMiddleware(jwtService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(authMiddleware)
	{
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}

	return r
}
