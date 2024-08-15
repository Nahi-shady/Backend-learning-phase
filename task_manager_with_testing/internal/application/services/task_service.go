package services

import (
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) CreateTask(title string, userID primitive.ObjectID) error {
	task := &entities.Task{
		Title:     title,
		Completed: false,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.taskRepo.CreateTask(task)
}

func (s *TaskService) GetTasks(userID primitive.ObjectID) ([]entities.Task, error) {
	return s.taskRepo.GetTasksByUserID(userID)
}

func (s *TaskService) UpdateTask(taskID primitive.ObjectID, title string, completed bool) error {
	return s.taskRepo.UpdateTask(taskID, title, completed)
}

func (s *TaskService) DeleteTask(taskID primitive.ObjectID) error {
	return s.taskRepo.DeleteTask(taskID)
}
