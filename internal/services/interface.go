// interface.go
//
// Interfaces for service methods
package services

import (
	"todo/internal/models"
)

// Service methods to work with tasks through interface
type Service interface {
	GetAllTasksService() ([]models.Task, error)
	AddTaskService(text string) error
	EditTaskService(index int, newText string) error
	DeleteTaskService(index int) error
	ToggleTaskService(index int) error
}
