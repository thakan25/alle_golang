package models

import "time"

// TaskStatus represents the possible states of a task
type TaskStatus string

const (
	TaskStatusPending     TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

// Task represents a task in our system
type Task struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	DueDate     Date       `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
} 