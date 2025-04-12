package dtos

import "time"

// UserDTO represents the user data transfer object for service layer
type UserDTO struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is not included in JSON responses
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserDTO represents the DTO for creating a user
type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserDTO represents the DTO for updating a user
type UpdateUserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
} 