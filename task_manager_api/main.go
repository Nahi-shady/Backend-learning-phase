package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	r := gin.Default()
	r.GET("/tasks", getTasks)
	r.GET("/task/:id", getTaskByID)
	r.PUT("/task/:id", updateTask)
	r.DELETE("/task/:id", deleteTask)
	r.POST("/task", createTask)

	r.Run()
}

func getTasks(c *gin.Context) {
	c.JSON(200, tasks)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(200, task)
			return
		}
	}
	c.JSON(404, gin.H{"error": "I ain't got nothin"})
}

func updateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "not found!"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(200, task)
			return
		}
	}
	c.JSON(404, gin.H{"error": "I ain't got nothin like that"})
}

func createTask(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}
