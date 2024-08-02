// data/task_service.go

package data

import (
	"errors"
	"sync"

	"task_manager/models"
)

var (
	tasks        = make(map[int64]*models.Task)
	nextID int64 = 1
	mutex  sync.Mutex
)

func GetTasks() []*models.Task {
	mutex.Lock()
	defer mutex.Unlock()

	taskList := make([]*models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	return taskList
}

func GetTaskByID(id int64) (*models.Task, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if task, exists := tasks[id]; exists {
		return task, nil
	}
	return nil, errors.New("task not found")
}

func CreateTask(task *models.Task) *models.Task {
	mutex.Lock()
	defer mutex.Unlock()

	task.ID = nextID
	nextID++
	tasks[task.ID] = task

	return task
}

func UpdateTask(id int64, updatedTask *models.Task) (*models.Task, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if task, exists := tasks[id]; exists {
		task.Title = updatedTask.Title
		task.Description = updatedTask.Description
		task.DueDate = updatedTask.DueDate
		task.Status = updatedTask.Status
		return task, nil
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id int64) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := tasks[id]; exists {
		delete(tasks, id)
		return nil
	}
	return errors.New("task not found")
}
