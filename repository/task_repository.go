package repository

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"task-manager/common"
	"task-manager/models"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id string) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	GetByStatus(status models.TaskStatus) ([]*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
}

// InMemoryTaskRepository implements TaskRepository using in-memory storage
type InMemoryTaskRepository struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

// NewInMemoryTaskRepository creates a new in-memory task repository
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*models.Task),
	}
}

func (r *InMemoryTaskRepository) Create(task *models.Task) error {
	if task == nil {
		return common.ErrInvalidTaskStatus
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) GetByID(id string) (*models.Task, error) {
	if id == "" {
		return nil, common.ErrInvalidTaskStatus
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, common.ErrTaskNotFound
	}
	return task, nil
}

func (r *InMemoryTaskRepository) GetAll() ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) GetByStatus(status models.TaskStatus) ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) Update(task *models.Task) error {
	if task == nil {
		return common.ErrInvalidTaskStatus
	}
	if task.ID == "" {
		return common.ErrInvalidTaskStatus
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return common.ErrTaskNotFound
	}

	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(id string) error {
	if id == "" {
		return common.ErrInvalidTaskStatus
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return common.ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
} 