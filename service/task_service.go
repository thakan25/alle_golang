package service

import (
	"errors"
	"time"

	"task-manager/models"
	"task-manager/repository"
)

var (
	ErrInvalidTaskStatus = errors.New("invalid task status")
)

// TaskService handles business logic for tasks
type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new task service
func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask creates a new task with business logic
func (s *TaskService) CreateTask(task *models.Task) error {
	// Validate task
	if err := s.validateTask(task); err != nil {
		return err
	}

	// Set timestamps
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	return s.repo.Create(task)
}

// GetTask retrieves a task by ID
func (s *TaskService) GetTask(id string) (*models.Task, error) {
	return s.repo.GetByID(id)
}

// GetTasks retrieves all tasks
func (s *TaskService) GetTasks() ([]*models.Task, error) {
	return s.repo.GetAll()
}

// UpdateTask updates an existing task
func (s *TaskService) UpdateTask(task *models.Task) error {
	// Validate task
	if err := s.validateTask(task); err != nil {
		return err
	}

	// Set update timestamp
	task.UpdatedAt = time.Now()

	return s.repo.Update(task)
}

// DeleteTask deletes a task by ID
func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}

// validateTask validates task fields
func (s *TaskService) validateTask(task *models.Task) error {
	// Set default status if not provided
	if task.Status == "" {
		task.Status = models.StatusPending
	}

	// Validate status
	switch task.Status {
	case models.StatusPending, models.StatusInProgress, models.StatusCompleted:
		// Valid status
	default:
		return ErrInvalidTaskStatus
	}

	return nil
} 