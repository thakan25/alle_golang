package errors

import (
	"encoding/json"
	"net/http"

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
		
		// Use defer to recover from panics
		defer func() {
			if err := recover(); err != nil {
				handleError(w, err.(error))
			}
		}()

		next.ServeHTTP(rw, r)
	})
}

// handleError handles specific errors and returns appropriate HTTP responses
func handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	switch err {
	case repository.ErrTaskNotFound:
		w.WriteHeader(http.StatusNotFound)
	case service.ErrInvalidTaskStatus:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

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