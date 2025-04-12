package service

import (
	"time"
	"strings"

	"task-manager/adapters"
	"task-manager/common"
	"task-manager/logging"
	"task-manager/models"
	"task-manager/models/dtos"
	"task-manager/repository"
	"github.com/google/uuid"
)

// TaskService handles business logic for tasks
type TaskService struct {
	repo    repository.TaskRepository
	userRepo repository.UserRepository
	adapter *adapters.ServiceToRepositoryAdapter
}

// NewTaskService creates a new task service
func NewTaskService(repo repository.TaskRepository, userRepo repository.UserRepository) *TaskService {
	logging.Info("Initializing new TaskService")
	return &TaskService{
		repo:     repo,
		userRepo: userRepo,
		adapter:  adapters.NewServiceToRepositoryAdapter(),
	}
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(dto dtos.CreateTaskDTO) (*dtos.TaskDTO, error) {
	logging.Info("Creating new task for user ID: %s", dto.UserID)
	logging.Info("Task details - Title: %s, Description: %s, DueDate: %v", dto.Title, dto.Description, dto.DueDate)

	// Validate user exists
	user, err := s.userRepo.GetByID(dto.UserID)
	if err != nil {
		logging.Error("Error finding user: %v", err)
		return nil, err
	}
	if user == nil {
		logging.Error("User not found with ID: %s", dto.UserID)
		return nil, common.ErrUserNotFound
	}
	logging.Info("Found user for task creation: %v", user)

	// Create task entity with formatted UUID (T prefix and no hyphens)
	uuidStr := uuid.New().String()
	formattedUUID := "T" + strings.ReplaceAll(uuidStr, "-", "")
	logging.Info("Generated task ID: %s", formattedUUID)

	// Create task DTO
	taskDTO := &dtos.TaskDTO{
		ID:          formattedUUID,
		UserID:      dto.UserID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      string(models.TaskStatusPending), // Always set as pending for new tasks
		DueDate:     dto.DueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	logging.Info("Created task DTO: %+v", taskDTO)

	// Convert to entity
	task := s.adapter.ToTaskEntity(*taskDTO)
	logging.Info("Converted to task entity: %+v", task)

	// Save task
	if err := s.repo.Create(task); err != nil {
		logging.Error("Error saving task: %v", err)
		return nil, err
	}
	logging.Info("Successfully saved task to repository")

	// Convert back to DTO
	resultDTO := s.adapter.ToTaskDTO(task)
	logging.Info("Task creation completed successfully. Final DTO: %+v", resultDTO)
	return &resultDTO, nil
}

// GetTasks retrieves all tasks
func (s *TaskService) GetTasks(status string) ([]*dtos.TaskDTO, error) {
	logging.Info("Retrieving all tasks with status: %s", status)
	tasks, err := s.repo.GetAll()
	if err != nil {
		logging.Error("Error retrieving tasks: %v", err)
		return nil, err
	}

	// Convert to DTOs
	taskDTOs := make([]*dtos.TaskDTO, len(tasks))
	for i, task := range tasks {
		dto := s.adapter.ToTaskDTO(task)
		taskDTOs[i] = &dto
	}

	// Filter by status if provided
	if status != "" {
		filteredDTOs := make([]*dtos.TaskDTO, 0)
		for _, dto := range taskDTOs {
			if dto.Status == status {
				filteredDTOs = append(filteredDTOs, dto)
			}
		}
		taskDTOs = filteredDTOs
	}

	logging.Info("Retrieved %d tasks", len(taskDTOs))
	return taskDTOs, nil
}

// GetTask retrieves a task by ID
func (s *TaskService) GetTask(id string) (*dtos.TaskDTO, error) {
	logging.Info("Fetching task with ID: %s", id)
	task, err := s.repo.GetByID(id)
	if err != nil {
		logging.Error("Failed to fetch task with ID %s: %v", id, err)
		return nil, err
	}
	logging.Info("Successfully fetched task with ID: %s", id)
	dto := s.adapter.ToTaskDTO(task)
	return &dto, nil
}

// UpdateTask updates an existing task
func (s *TaskService) UpdateTask(dto dtos.UpdateTaskDTO) (*dtos.TaskDTO, error) {
	logging.Info("Updating task with ID: %s", dto.ID)

	// Create TaskDTO from UpdateTaskDTO
	taskDTO := dtos.TaskDTO{
		ID:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		DueDate:     dto.DueDate,
		UpdatedAt:   time.Now(),
	}

	// Get existing task to preserve CreatedAt
	existingTask, err := s.repo.GetByID(dto.ID)
	if err != nil {
		logging.Error("Failed to get existing task: %v", err)
		return nil, err
	}
	taskDTO.CreatedAt = existingTask.CreatedAt

	// Convert to entity
	task := s.adapter.ToTaskEntity(taskDTO)

	// Validate task
	if err := s.validateTask(task); err != nil {
		logging.Error("Task validation failed: %v", err)
		return nil, err
	}

	if err := s.repo.Update(task); err != nil {
		logging.Error("Failed to update task in repository: %v", err)
		return nil, err
	}

	// Convert back to DTO
	resultDTO := s.adapter.ToTaskDTO(task)
	logging.Info("Task updated successfully with ID: %s", dto.ID)
	return &resultDTO, nil
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
	if task.UserID == "" {
		return common.ErrInvalidTaskStatus
	}
	if task.Title == "" {
		return common.ErrInvalidTaskStatus
	}
	if !s.isValidStatus(task.Status) {
		return common.ErrInvalidTaskStatus
	}
	if task.DueDate.IsZero() {
		return common.ErrInvalidTaskStatus
	}
	return nil
}

// isValidStatus checks if the given status is valid
func (s *TaskService) isValidStatus(status models.TaskStatus) bool {
	return status == models.TaskStatusPending || 
	       status == models.TaskStatusInProgress || 
	       status == models.TaskStatusCompleted
}

// validateTaskStatus validates the task status
func (s *TaskService) validateTaskStatus(status string) error {
	if status != string(models.TaskStatusPending) &&
	   status != string(models.TaskStatusInProgress) &&
	   status != string(models.TaskStatusCompleted) {
		return common.ErrInvalidTaskStatus
	}
	return nil
}

// GetTasksByUserID retrieves all tasks for a specific user
func (s *TaskService) GetTasksByUserID(userID string) ([]*dtos.TaskDTO, error) {
	logging.Info("Retrieving tasks for user ID: %s", userID)

	// Validate user exists
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		logging.Error("Error finding user: %v", err)
		return nil, err
	}
	if user == nil {
		logging.Error("User not found with ID: %s", userID)
		return nil, common.ErrUserNotFound
	}
	logging.Info("Found user: %v", user)

	// Get all tasks
	tasks, err := s.repo.GetAll()
	if err != nil {
		logging.Error("Error retrieving tasks: %v", err)
		return nil, err
	}
	logging.Info("Retrieved %d total tasks", len(tasks))

	// Filter tasks for the user
	var userTasks []*dtos.TaskDTO
	for _, task := range tasks {
		if task.UserID == userID {
			dto := s.adapter.ToTaskDTO(task)
			userTasks = append(userTasks, &dto)
		}
	}
	logging.Info("Found %d tasks for user", len(userTasks))
	return userTasks, nil
} 