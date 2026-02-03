// main.go
//
// Entrypoint for interactive mode of app
package main

import (
	"todo/internal/repositories"
	"todo/internal/services"
	"todo/internal/ui"
)

// main: entrypoint function
func main() {
	repo := repositories.NewTaskRepository()
	service := services.NewTaskService(repo)
	menu := ui.CreateMenu(service)

	menu()
}
