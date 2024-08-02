package models

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"` // Format: YYYY-MM-DD
	Status      string `json:"status"`   // Example statuses: pending, completed
}
