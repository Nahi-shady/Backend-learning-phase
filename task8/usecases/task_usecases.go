package usecases

import (
	"task-manager/domain"
	"task-manager/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	taskRepo repositories.TaskRepository
}

func NewTaskUsecase(taskRepo repositories.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		taskRepo: taskRepo,
	}
}

func (t *TaskUsecase) CreateTask(title string, userID primitive.ObjectID) error {
	task := &domain.Task{
		Title:     title,
		Completed: false,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return t.taskRepo.CreateTask(task)
}

func (t *TaskUsecase) GetTasks(userID primitive.ObjectID) ([]domain.Task, error) {
	return t.taskRepo.GetTasksByUserID(userID)
}

func (t *TaskUsecase) UpdateTask(taskID primitive.ObjectID, title string, completed bool) error {
	return t.taskRepo.UpdateTask(taskID, title, completed)
}

func (t *TaskUsecase) DeleteTask(taskID primitive.ObjectID) error {
	return t.taskRepo.DeleteTask(taskID)
}
