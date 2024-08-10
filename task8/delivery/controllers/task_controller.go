package controllers

import (
	"net/http"

	"task-manager/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskUsecase *usecases.TaskUsecase
}

func NewTaskController(taskUsecase *usecases.TaskUsecase) *TaskController {
	return &TaskController{
		taskUsecase: taskUsecase,
	}
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
		if err := c.taskUsecase.CreateTask(req.Title, id); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Task created"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	if userIDStr, ok := userID.(string); ok {
		id, _ := primitive.ObjectIDFromHex(userIDStr)
		tasks, err := c.taskUsecase.GetTasks(id)
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

	if err := c.taskUsecase.UpdateTask(taskID, req.Title, req.Completed); err != nil {
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

	if err := c.taskUsecase.DeleteTask(taskID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}