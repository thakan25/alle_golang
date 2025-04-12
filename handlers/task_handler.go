package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"task-manager/adapters"
	"task-manager/logging"
	"task-manager/models/requests"
	"task-manager/service"
)

type TaskHandler struct {
	service *service.TaskService
	adapter *adapters.ControllerToServiceAdapter
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
		adapter: adapters.NewControllerToServiceAdapter(),
	}
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
	var req requests.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logging.Error("Failed to decode task request body: %v", err)
		panic(err)
	}

	// Convert request to service DTO
	createTaskDto := h.adapter.ToCreateTaskDTO(req)

	// Call service
	taskDTO, err := h.service.CreateTask(createTaskDto)
	if err != nil {
		logging.Error("Failed to create task: %v", err)
		panic(err)
	}

	// Convert service DTO to response
	resp := h.adapter.ToTaskResponse(*taskDTO)
	h.respondJSON(w, http.StatusCreated, resp)
}

// GetTasks handles getting all tasks
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	statusParam := r.URL.Query().Get("status")
	
	taskDTOs, err := h.service.GetTasks(statusParam)
	if err != nil {
		logging.Error("Error getting tasks: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to responses
	responses := h.adapter.ToTaskResponses(taskDTOs)
	json.NewEncoder(w).Encode(responses)
}

// GetTasksByUserID handles getting tasks for a specific user
func (h *TaskHandler) GetTasksByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	tasks, err := h.service.GetTasksByUserID(userID)
	if err != nil {
		logging.Error("Error getting tasks for user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to responses
	responses := h.adapter.ToTaskResponses(tasks)
	json.NewEncoder(w).Encode(responses)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	taskDTO, err := h.service.GetTask(id)
	if err != nil {
		logging.Error("Failed to get task with ID %s: %v", id, err)
		panic(err)
	}

	// Convert service DTO to response
	resp := h.adapter.ToTaskResponse(*taskDTO)
	h.respondJSON(w, http.StatusOK, resp)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req requests.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logging.Error("Failed to decode task update request body: %v", err)
		panic(err)
	}

	// Convert request to service DTO
	updateDTO := h.adapter.ToUpdateTaskDTO(id, req)

	// Call service
	taskDTO, err := h.service.UpdateTask(updateDTO)
	if err != nil {
		logging.Error("Failed to update task with ID %s: %v", id, err)
		panic(err)
	}

	// Convert service DTO to response
	resp := h.adapter.ToTaskResponse(*taskDTO)
	h.respondJSON(w, http.StatusOK, resp)
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