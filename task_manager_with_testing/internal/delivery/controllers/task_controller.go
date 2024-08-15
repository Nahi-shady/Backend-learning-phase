package controllers

import (
	"net/http"
	"task-manager/internal/application/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID, _ := ctx.Get("userID")
	if userIDStr, ok := userID.(string); ok {
		id, _ := primitive.ObjectIDFromHex(userIDStr)
		if err := c.taskService.CreateTask(req.Title, id); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Task created", "ID: ": userID})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	if userIDStr, ok := userID.(string); ok {
		id, _ := primitive.ObjectIDFromHex(userIDStr)
		tasks, err := c.taskService.GetTasks(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	var req struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	taskID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := c.taskService.UpdateTask(taskID, req.Title, req.Completed); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	taskID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := c.taskService.DeleteTask(taskID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
