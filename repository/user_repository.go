package repository

import (
	"task-manager/common"
	"task-manager/logging"
	"task-manager/models/dtos"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *dtos.UserDTO) error
	GetByID(id string) (*dtos.UserDTO, error)
	GetByEmail(email string) (*dtos.UserDTO, error)
	GetAll() ([]*dtos.UserDTO, error)
	Delete(id string) error
}

// InMemoryUserRepository implements UserRepository using in-memory storage
type InMemoryUserRepository struct {
	users map[string]*dtos.UserDTO
}

// NewInMemoryUserRepository creates a new in-memory user repository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*dtos.UserDTO),
	}
}

// Create adds a new user to the repository
func (r *InMemoryUserRepository) Create(user *dtos.UserDTO) error {
	logging.Info("Creating user with ID: %s", user.ID)
	r.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by ID
func (r *InMemoryUserRepository) GetByID(id string) (*dtos.UserDTO, error) {
	logging.Info("Looking up user with ID: %s", id)
	logging.Info("Current users in repository: %v", r.users)
	user, exists := r.users[id]
	if !exists {
		logging.Error("User not found with ID: %s", id)
		return nil, common.ErrUserNotFound
	}
	logging.Info("Found user: %v", user)
	return user, nil
}

// GetByEmail retrieves a user by email
func (r *InMemoryUserRepository) GetByEmail(email string) (*dtos.UserDTO, error) {
	logging.Info("Looking up user with email: %s", email)
	for _, user := range r.users {
		if user.Email == email {
			logging.Info("Found user: %v", user)
			return user, nil
		}
	}
	logging.Error("User not found with email: %s", email)
	return nil, common.ErrUserNotFound
}

// GetAll retrieves all users
func (r *InMemoryUserRepository) GetAll() ([]*dtos.UserDTO, error) {
	users := make([]*dtos.UserDTO, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	logging.Info("Retrieved %d users", len(users))
	return users, nil
}

// Delete removes a user by ID
func (r *InMemoryUserRepository) Delete(id string) error {
	if _, exists := r.users[id]; !exists {
		return common.ErrUserNotFound
	}
	delete(r.users, id)
	return nil
} 