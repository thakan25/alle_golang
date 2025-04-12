package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"task-manager/config"
	"task-manager/handlers"
	"task-manager/repository"
	"task-manager/service"
	"task-manager/errors"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize repositories
	taskRepo := repository.NewInMemoryTaskRepository()
	userRepo := repository.NewInMemoryUserRepository()

	// Initialize services
	taskService := service.NewTaskService(taskRepo, userRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize router
	router := mux.NewRouter()

	// Apply error handling middleware
	router.Use(errors.ErrorHandler)

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// User endpoints
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Task endpoints
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	router.HandleFunc("/users/{userId}/tasks", taskHandler.GetTasksByUserID).Methods("GET")

	// Start server
	log.Printf("Server starting on port %d", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.ServerPort), router))
} 