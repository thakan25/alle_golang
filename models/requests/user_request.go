package requests

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
} 