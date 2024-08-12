// data/task_service.go

package data

import (
	"errors"
	"time"

	"task_manager/models"
)

var (
	tasks        = make(map[int64]*models.Task)
	nextID int64 = 1
)

func GetTasks() []*models.Task {
	taskList := make([]*models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	return taskList
}

func GetTaskByID(id int64) (*models.Task, error) {
	if task, exists := tasks[id]; exists {
		return task, nil
	}
	return nil, errors.New("task not found")
}

func CreateTask(task *models.Task) *models.Task {
	task.ID = nextID
	nextID++

	task.DueDate = time.Now()
	task.Status = "pending"
	tasks[task.ID] = task

	return task
}

func UpdateTask(id int64, updatedTask *models.Task) (*models.Task, error) {
	if task, exists := tasks[id]; exists {
		if updatedTask.Title != "" {
			task.Title = updatedTask.Title
		}
		if updatedTask.Description != "" {
			task.Description = updatedTask.Description
		}
		if updatedTask.Status != "" {
			task.Status = updatedTask.Status
		}
		return task, nil
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id int64) error {
	if _, exists := tasks[id]; exists {
		delete(tasks, id)
		return nil
	}
	return errors.New("task not found")
}
