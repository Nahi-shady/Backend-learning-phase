package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasksHandler)
	router.GET("/tasks/:id", controllers.GetTaskByIDHandler)
	router.POST("/tasks", controllers.CreateTaskHandler)
	router.PUT("/tasks/:id", controllers.UpdateTaskHandler)
	router.DELETE("/tasks/:id", controllers.DeleteTaskHandler)

	return router
}
