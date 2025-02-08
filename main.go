package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-console-app/service"
)

func main() {
	//TODO: Implement UpdateTask  []
	//TODO: Implement DeleteTask  []
	terminator := false
	reader := bufio.NewReader(os.Stdin)
	var choice int
	for terminator != true {
		println("Select an option")
		println("Select 1 to get task by id")
		println("Select 2 to get all tasks")
		println("Select 3 to create task")
		println("Select 4 to update a task")
		println("Select 5 to delete a task")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var id string
			fmt.Print("Type task Id:")
			fmt.Scanln(&id)
			response := service.GetTaskById(id)
			println(response)
		case 2:
			println("All available tasks:")
			response := service.GetAllTasks()
			fmt.Println(response)
		case 3:
			fmt.Print("Type task title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			response := service.CreateTask(title)
			fmt.Println(response)
		case 4:
			println("updating task...")
		case 5:
			println("deleting task...")
		default:
			println("option not available...")
		}
		var answer string
		fmt.Println("Wanna Continue? Y/N")
		fmt.Scanln(&answer)

		if strings.ToUpper(answer) == "N" {
			terminator = true
		}

	}

}
