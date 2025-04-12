package responses

import (
	"time"
	"task-manager/models"
)

// TaskResponse represents the response body for task operations
type TaskResponse struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     models.Date `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// TasksResponse represents the response body for multiple tasks
type TasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
} 