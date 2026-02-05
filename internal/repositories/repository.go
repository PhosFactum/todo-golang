// repository.go
//
// Layer to take data and give it to service
package repositories

import (
	"todo/internal/models"
	"errors"
)

// TaskRepository: storage for data from models (slice of tasks now)
type TaskRepository struct {
	tasks []models.Task	
}

// NewTaskRepository: constructor for creating storage for tasks
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: []models.Task{},
	}
}

// GetAll: returns slice of tasks
func (r *TaskRepository) GetAll() ([]models.Task, error) {
	return r.tasks, nil
}

// Create: creates new task in data layer
func (r *TaskRepository) Create(task models.Task) error {
	r.tasks = append(r.tasks, task)
	return nil
}

// Edit: edits task by index via changing text
func (r *TaskRepository) Edit(index int, newText string) error {
	if index < 0 || index >= len(r.tasks) {
		return errors.New("Invalid index!")
	}

	r.tasks[index].Text = newText
	return nil
}

// Delete: deletes task by index
func (r *TaskRepository) Delete(index int) error {
	if index < 0 || index >= len(r.tasks) {
		return errors.New("Invalid index!")
	}

	r.tasks = append(r.tasks[:index], r.tasks[index+1:]...)
	return nil
}

// Toggle: changes the status 'done' to 'undone' and vice versa
func (r *TaskRepository) Toggle(index int) error {
	if index < 0 || index >= len(r.tasks) {
		return errors.New("Invalid index!")
	}

	r.tasks[index].Done = !r.tasks[index].Done
	return nil
}
