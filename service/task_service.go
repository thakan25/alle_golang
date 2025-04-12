package service

import (
	"errors"
	"time"

	"task-manager/logging"
	"task-manager/models"
	"task-manager/repository"
)

var (
	ErrInvalidTaskStatus = errors.New("invalid task status")
	ErrInvalidTaskTitle  = errors.New("invalid task title")
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
	logging.Info("Creating new task with title: %s", task.Title)

	// Set default status if not provided
	if task.Status == "" {
		task.Status = models.TaskStatusPending
	}

	// Validate task
	if err := s.validateTask(task); err != nil {
		logging.Error("Task validation failed: %v", err)
		return err
	}

	// Set timestamps
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	if err := s.repo.Create(task); err != nil {
		logging.Error("Failed to create task in repository: %v", err)
		return err
	}

	logging.Info("Task created successfully with ID: %s", task.ID)
	return nil
}

// GetTask retrieves a task by ID
func (s *TaskService) GetTask(id string) (*models.Task, error) {
	logging.Info("Fetching task with ID: %s", id)
	task, err := s.repo.GetByID(id)
	if err != nil {
		logging.Error("Failed to fetch task with ID %s: %v", id, err)
		return nil, err
	}
	logging.Info("Successfully fetched task with ID: %s", id)
	return task, nil
}

// GetTasks retrieves all tasks or filters by status if provided
func (s *TaskService) GetTasks(statusParam string) ([]*models.Task, error) {
	if statusParam != "" {
		// Validate status
		status := models.TaskStatus(statusParam)
		if !s.isValidStatus(status) {
			logging.Error("Invalid status parameter: %s", statusParam)
			return nil, ErrInvalidTaskStatus
		}

		logging.Info("Fetching tasks with status: %s", status)
		tasks, err := s.repo.GetByStatus(status)
		if err != nil {
			logging.Error("Failed to fetch tasks with status %s: %v", status, err)
			return nil, err
		}
		logging.Info("Successfully fetched %d tasks with status %s", len(tasks), status)
		return tasks, nil
	}

	logging.Info("Fetching all tasks")
	tasks, err := s.repo.GetAll()
	if err != nil {
		logging.Error("Failed to fetch tasks from repository: %v", err)
		return nil, err
	}
	logging.Info("Successfully fetched %d tasks", len(tasks))
	return tasks, nil
}

// isValidStatus checks if the given status is valid
func (s *TaskService) isValidStatus(status models.TaskStatus) bool {
	return status == models.TaskStatusPending || 
	       status == models.TaskStatusInProgress || 
	       status == models.TaskStatusCompleted
}

// UpdateTask updates an existing task
func (s *TaskService) UpdateTask(task *models.Task) error {
	logging.Info("Updating task with ID: %s", task.ID)

	// Validate task
	if err := s.validateTask(task); err != nil {
		logging.Error("Task validation failed: %v", err)
		return err
	}

	// Update timestamp
	task.UpdatedAt = time.Now()

	if err := s.repo.Update(task); err != nil {
		logging.Error("Failed to update task in repository: %v", err)
		return err
	}

	logging.Info("Task updated successfully with ID: %s", task.ID)
	return nil
}

// DeleteTask deletes a task by ID
func (s *TaskService) DeleteTask(id string) error {
	logging.Info("Deleting task with ID: %s", id)
	if err := s.repo.Delete(id); err != nil {
		logging.Error("Failed to delete task with ID %s: %v", id, err)
		return err
	}
	logging.Info("Task deleted successfully with ID: %s", id)
	return nil
}

// validateTask validates task fields
func (s *TaskService) validateTask(task *models.Task) error {
	if task.Title == "" {
		return ErrInvalidTaskTitle
	}
	if task.Status != models.TaskStatusPending && 
	   task.Status != models.TaskStatusInProgress && 
	   task.Status != models.TaskStatusCompleted {
		return ErrInvalidTaskStatus
	}
	return nil
} 