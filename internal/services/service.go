// service.go
//
// All services for Add, Edit, Show, Delete and Mark operations with tasks
package services

import (
	"todo/internal/repositories"
	"todo/internal/models"
)

// TaskService: structure for link to our repository (storage)
type TaskService struct {
	repo *repositories.TaskRepository
}

// NewTaskService: constructor for create service with link on repo
func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// GetAllService: reads tasks from data layer
func (s *TaskService) GetAllTasksService() ([]models.Task, error) {
	return s.repo.GetAll()
}

// AddTaskService: adding operation for creating new task
func (s *TaskService) AddTaskService(text string) error {
	// creating new task object
	task := models.Task{
		Text: text,
		Done: false,
	}

	s.repo.Create(task)
	return nil	
}


// EditTask: edits existing task
func (s *TaskService) EditTaskService(index int, newText string) error {
	return s.repo.Edit(index, newText)
}

// DeleteTask: deletes existing task
func (s *TaskService) DeleteTaskService(index int) error {
	return s.repo.Delete(index)
}

// ToggleTask: toggle status of completion on task
func (s *TaskService) ToggleTaskService(index int) error {
	return s.repo.Toggle(index)
}
