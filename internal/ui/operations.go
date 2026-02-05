// operations.go 
//
// Need to call functions from 'services' package like a wrap
package ui

import (
	"todo/internal/services"
	"fmt"
)

// ShowTasks: shows all tasks 
func ShowTasks(service *services.TaskService) {
	fmt.Println("=== Tasks list ===")

	tasks, err := service.GetAllTasksService()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}

	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}
}

// AddTask: adds new task
func AddTask(service *services.TaskService) {
	fmt.Println("=== Creating new task ===")

	fmt.Print("Write your task: ")
	text, err := GetString()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = service.AddTaskService(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// EditTask: edits existing task
func EditTask(service *services.TaskService) {
	fmt.Println("=== Editing old task ===")
	
	fmt.Print("Type index of task do edit: ")
	i, err := GetInt()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	i--

	fmt.Print("Write new text for this task: ")
	editedText, err := GetString()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = service.EditTaskService(i, editedText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task edited successfully!")
}

// DeleteTask: deletes existing task
func DeleteTask(service *services.TaskService) {
	fmt.Println("=== Deleting old task ===")

	fmt.Print("Type index of task do delete: ")
	i, err := GetInt()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	i--

	err = service.DeleteTaskService(i)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task deleted successfully!")
}

// ToggleTask: mark task as done / undone
func ToggleTask(service *services.TaskService) {
	fmt.Println("=== Task status toggler ===")

	fmt.Print("Type index of task to toggle: ")
	i, err := GetInt()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	i--

	err = service.ToggleTaskService(i)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task toggled successfully!")
}


/* TODO:
- Edit: doesn't work
- Toggle: doesn't work

*/
