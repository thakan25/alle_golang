package requests

import (
	"task-manager/models"
)

// CreateTaskRequest represents the request body for creating a task
type CreateTaskRequest struct {
	UserID      string     `json:"user_id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	DueDate     models.Date `json:"due_date" binding:"required"`
}

// UpdateTaskRequest represents the request body for updating a task
type UpdateTaskRequest struct {
	UserID      string     `json:"user_id"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     models.Date `json:"due_date"`
} 