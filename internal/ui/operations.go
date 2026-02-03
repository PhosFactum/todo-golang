// operations.go 
//
// Need to call functions from 'services' package like a wrap
package ui

import (
	"todo/internal/services"
	"log"
	"fmt"
)

// ShowTasks: shows all tasks 
func ShowTasks(service *services.TaskService) {
	fmt.Println("=== Tasks list ===")

	_, err := service.GetAllTasksService()
	if err != nil {
		log.Fatal(err)
	}
}

// AddTask: adds new task
func AddTask(service *services.TaskService) {
	fmt.Println("=== Creating new task ===")

	fmt.Print("Write your task: ")
	text, err := GetString()
	if err != nil {
		log.Fatal(err)
	}

	err = service.AddTaskService(text)
	if err != nil {
		log.Fatal(err)
	}
}

// EditTask: edits existing task
func EditTask(service *services.TaskService) {
	fmt.Println("=== Editing old task ===")

	// USER INPUT LOGIC WILL BE HERE

	editedText := "editedText"
	i := 2
	//

	err := service.EditTaskService(i, editedText)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteTask: deletes existing task
func DeleteTask(service *services.TaskService) {
	fmt.Println("=== Deleting old task ===")

	// USER INPUT LOGIC WILL BE HERE

	i := 1
	//

	err := service.DeleteTaskService(i)
	if err != nil {
		log.Fatal(err)
	}
}

// ToggleTask: mark task as done / undone
func ToggleTask(service *services.TaskService) {
	fmt.Println("=== Task status toggler ===")

	// USER INPUT LOGIC WILL BE HERE
	
	i := 2
	//

	err := service.ToggleTaskService(i)
	if err != nil {
		log.Fatal(err)
	}
}


/* TODO:
+- Show: doesn't show anything, but works
- Edit: doesn't work
+ Add: works, probably
+ Delete: doesn't receive index, but works
- Toggle: doesn't work

*/
