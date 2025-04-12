package errors

import (
	"encoding/json"
	"net/http"

	"task-manager/logging"
	"task-manager/repository"
	"task-manager/service"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorHandler is a middleware that handles errors
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer to capture the response
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		// Log request details
		logging.Info("Incoming request: %s %s", r.Method, r.URL.Path)
		
		// Use defer to recover from panics
		defer func() {
			if err := recover(); err != nil {
				logging.Error("Panic recovered: %v", err)
				handleError(w, err.(error))
			}
		}()

		next.ServeHTTP(rw, r)
	})
}

// handleError handles specific errors and returns appropriate HTTP responses
func handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	var statusCode int
	switch err {
	case repository.ErrTaskNotFound:
		statusCode = http.StatusNotFound
		logging.Error("Task not found: %v", err)
	case service.ErrInvalidTaskStatus:
		statusCode = http.StatusBadRequest
		logging.Error("Invalid task status: %v", err)
	default:
		statusCode = http.StatusInternalServerError
		logging.Error("Internal server error: %v", err)
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

// Custom response writer to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
} 