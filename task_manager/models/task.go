package models

import (
	"time"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"` // Format: YYYY-MM-DD
	Status      string    `json:"status"`   // Example statuses: pending, completed
}
