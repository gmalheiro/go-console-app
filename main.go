package main

import (
	"fmt"
	"go-console-app/controller"
	"strings"
)

func main() {
	taskController := controller.NewTaskController()

	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1 - Get task by id")
		fmt.Println("2 - Get all tasks")
		fmt.Println("3 - Create task")
		fmt.Println("4 - Update a task")
		fmt.Println("5 - Delete a task")
		fmt.Println("6 - Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			taskController.GetTaskById()
		case 2:
			taskController.GetAllTasks()
		case 3:
			taskController.CreateTask()
		case 4:
			taskController.UpdateTask()
		case 5:
			taskController.DeleteTask()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, try again.")
		}

		var answer string
		fmt.Print("\nWanna Continue? (Y/N): ")
		fmt.Scanln(&answer)
		if strings.ToUpper(answer) == "N" {
			break
		}
	}
}
