package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"task-manager/logging"
	"task-manager/models"
	"task-manager/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// respondJSON writes a JSON response with the given status code
func (h *TaskHandler) respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		logging.Error("Failed to decode task request body: %v", err)
		panic(err)
	}

	if err := h.service.CreateTask(&task); err != nil {
		logging.Error("Failed to create task: %v", err)
		panic(err)
	}

	h.respondJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Get status from query parameter and pass it to service
	statusParam := r.URL.Query().Get("status")
	tasks, err := h.service.GetTasks(statusParam)
	if err != nil {
		logging.Error("Failed to get tasks: %v", err)
		panic(err)
	}

	h.respondJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task, err := h.service.GetTask(id)
	if err != nil {
		logging.Error("Failed to get task with ID %s: %v", id, err)
		panic(err)
	}

	h.respondJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		logging.Error("Failed to decode task update request body: %v", err)
		panic(err)
	}

	task.ID = id
	if err := h.service.UpdateTask(&task); err != nil {
		logging.Error("Failed to update task with ID %s: %v", id, err)
		panic(err)
	}

	h.respondJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteTask(id); err != nil {
		logging.Error("Failed to delete task with ID %s: %v", id, err)
		panic(err)
	}

	h.respondJSON(w, http.StatusNoContent, nil)
} 