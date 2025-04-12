package service

import (
	"errors"
	"strings"
	"task-manager/common"
	"task-manager/models/dtos"
	"task-manager/repository"
	"task-manager/utils"
	"time"

	"github.com/google/uuid"
)

// UserService handles business logic for users
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(dto dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	// Validate input
	if err := utils.ValidateEmail(dto.Email); err != nil {
		return nil, err
	}
	if err := utils.ValidatePassword(dto.Password); err != nil {
		return nil, err
	}
	if err := utils.ValidateUsername(dto.Username); err != nil {
		return nil, err
	}

	// Check if user already exists
	existingUser, err := s.repo.GetByEmail(dto.Email)
	if err != nil && !errors.Is(err, common.ErrUserNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, common.ErrEmailExists
	}

	// Create user entity with formatted UUID (U prefix and no hyphens)
	uuidStr := uuid.New().String()
	formattedUUID := "U" + strings.ReplaceAll(uuidStr, "-", "")

	// Create user entity with generated UUID
	user := &dtos.UserDTO{
		ID:        formattedUUID,
		Email:     dto.Email,
		Username:  dto.Username,
		Password:  dto.Password, // Note: Password should be hashed before storage
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save user
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id string) (*dtos.UserDTO, error) {
	return s.repo.GetByID(id)
}

// GetUsers retrieves all users
func (s *UserService) GetUsers() ([]*dtos.UserDTO, error) {
	return s.repo.GetAll()
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
} 