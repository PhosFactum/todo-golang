// menu.go
//
// Prints menu for interactive mode of app
package ui

import (
	"todo/internal/services"
	"fmt"
)

// CreateMenu: printing menu for interactive mode
func CreateMenu(service *services.TaskService) func() {
	return func () {
		fmt.Println("=== Welcome to GoToDo! ===")


		for {
			fmt.Println("----------------------")
			fmt.Println("| 0. Exit            |")
			fmt.Println("| 1. Show all tasks  |")
			fmt.Println("| 2. Add new task    |")
			fmt.Println("| 3. Edit task       |")
			fmt.Println("| 4. Delete task     |")
			fmt.Println("| 5. Toggle task     |")
			fmt.Println("----------------------")

			fmt.Print("Select menu item: ")
			choice, err := GetInt()
			if err != nil {
				fmt.Printf("Input error: %w.\n", err)
				fmt.Println("Please, try again.")
				continue
			}

			if choice < 0 || choice > 5 {
				fmt.Println("Invalid choice, try again.")
				continue
			}

			switch(choice) {
				case 0:
					fmt.Println("--- Program was completed ---")
					return
				case 1:
					ShowTasks(service)
				case 2:
					AddTask(service)
				case 3:
					EditTask(service)
				case 4:
					DeleteTask(service)
				case 5:
					ToggleTask(service)
				default:
					fmt.Println("Invalid choice, try again.")
			}

		}
	}
}
