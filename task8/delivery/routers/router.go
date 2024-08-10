package routers

import (
	"task-manager/delivery/controllers"
	infrastructures "task-manager/infrastructures"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, taskController *controllers.TaskController, jwtService infrastructures.JWTService) *gin.Engine {
	r := gin.Default()

	authMiddleware := infrastructures.AuthMiddleware(jwtService)

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
