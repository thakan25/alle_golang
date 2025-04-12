package responses

import "time"

// UserResponse represents the response body for user operations
type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UsersResponse represents the response body for multiple users
type UsersResponse struct {
	Users []UserResponse `json:"users"`
} 