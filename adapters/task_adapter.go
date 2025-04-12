package adapters

import (
	"task-manager/models"
	"task-manager/models/dtos"
	"task-manager/models/requests"
	"task-manager/models/responses"
)

// ControllerToServiceAdapter converts controller request to service DTO
type ControllerToServiceAdapter struct{}

func NewControllerToServiceAdapter() *ControllerToServiceAdapter {
	return &ControllerToServiceAdapter{}
}

func (a *ControllerToServiceAdapter) ToCreateTaskDTO(req requests.CreateTaskRequest) dtos.CreateTaskDTO {
	return dtos.CreateTaskDTO{
		Title:       req.Title,
		Description: req.Description,
		UserID:      req.UserID,
		DueDate:     req.DueDate,
	}
}

func (a *ControllerToServiceAdapter) ToUpdateTaskDTO(id string, req requests.UpdateTaskRequest) dtos.UpdateTaskDTO {
	return dtos.UpdateTaskDTO{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     req.DueDate,
	}
}

func (a *ControllerToServiceAdapter) ToTaskResponse(dto dtos.TaskDTO) responses.TaskResponse {
	return responses.TaskResponse{
		ID:          dto.ID,
		UserID:      dto.UserID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		DueDate:     dto.DueDate,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

// ToTasksResponse converts a slice of TaskDTOs to TasksResponse
func (a *ControllerToServiceAdapter) ToTasksResponse(dtos []*dtos.TaskDTO) responses.TasksResponse {
	tasks := make([]responses.TaskResponse, len(dtos))
	for i, dto := range dtos {
		tasks[i] = a.ToTaskResponse(*dto)
	}
	return responses.TasksResponse{Tasks: tasks}
}

func (a *ControllerToServiceAdapter) ToCreateUserDTO(req requests.CreateUserRequest) dtos.CreateUserDTO {
	return dtos.CreateUserDTO{
		Username: req.Username,
		Email:    req.Email,
	}
}

func (a *ControllerToServiceAdapter) ToUpdateUserDTO(id string, req requests.UpdateUserRequest) dtos.UpdateUserDTO {
	return dtos.UpdateUserDTO{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
	}
}

func (a *ControllerToServiceAdapter) ToUserResponse(dto dtos.UserDTO) responses.UserResponse {
	return responses.UserResponse{
		ID:        dto.ID,
		Username:  dto.Username,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func (a *ControllerToServiceAdapter) ToUsersResponse(dtos []dtos.UserDTO) responses.UsersResponse {
	users := make([]responses.UserResponse, len(dtos))
	for i, dto := range dtos {
		users[i] = a.ToUserResponse(dto)
	}
	return responses.UsersResponse{Users: users}
}

// ServiceToRepositoryAdapter converts service DTO to repository entity
type ServiceToRepositoryAdapter struct{}

func NewServiceToRepositoryAdapter() *ServiceToRepositoryAdapter {
	return &ServiceToRepositoryAdapter{}
}

func (a *ServiceToRepositoryAdapter) ToTaskEntity(dto dtos.TaskDTO) *models.Task {
	return &models.Task{
		ID:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      models.TaskStatus(dto.Status),
		DueDate:     dto.DueDate,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

func (a *ServiceToRepositoryAdapter) ToTaskDTO(entity *models.Task) dtos.TaskDTO {
	return dtos.TaskDTO{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Status:      string(entity.Status),
		DueDate:     entity.DueDate,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (a *ServiceToRepositoryAdapter) ToTaskEntities(dtos []dtos.TaskDTO) []*models.Task {
	entities := make([]*models.Task, len(dtos))
	for i, dto := range dtos {
		entities[i] = a.ToTaskEntity(dto)
	}
	return entities
}

func (a *ServiceToRepositoryAdapter) ToTaskDTOs(entities []*models.Task) []dtos.TaskDTO {
	dtos := make([]dtos.TaskDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = a.ToTaskDTO(entity)
	}
	return dtos
}

func (a *ServiceToRepositoryAdapter) ToUserEntity(dto dtos.UserDTO) *models.User {
	return &models.User{
		ID:        dto.ID,
		Username:  dto.Username,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func (a *ServiceToRepositoryAdapter) ToUserDTO(entity *models.User) dtos.UserDTO {
	return dtos.UserDTO{
		ID:        entity.ID,
		Username:  entity.Username,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (a *ServiceToRepositoryAdapter) ToUserEntities(dtos []dtos.UserDTO) []*models.User {
	entities := make([]*models.User, len(dtos))
	for i, dto := range dtos {
		entities[i] = a.ToUserEntity(dto)
	}
	return entities
}

func (a *ServiceToRepositoryAdapter) ToUserDTOs(entities []*models.User) []dtos.UserDTO {
	dtos := make([]dtos.UserDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = a.ToUserDTO(entity)
	}
	return dtos
} 