// interface.go
//
// Methods for Repository interface
package repositories

import (
	"todo/internal/models"
)

// Repository interface: contract of methods for repository
type Repository interface {
	GetAll() ([]models.Task, error)
	Create(task models.Task) error
	Edit(index int, newText string) error
	Delete(index int) error
	Toggle(index int) error
}
