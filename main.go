package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-console-app/service"
	"go-console-app/utils"
)

func main() {
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
			fmt.Println(utils.TaskToJsonString(*response))
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
			var id string
			fmt.Print("Type task Id:")
			fmt.Scanln(&id)
			task := service.GetTaskById(id)
			var answer string
			fmt.Print("Wanna edit both status and title? Y/N")
			fmt.Scanln(&answer)
			if strings.ToUpper(answer) != "Y" {
				fmt.Print("Type 1 to update task title and 2 to update task status: ")
				var editChoice int
				fmt.Scanln(&editChoice)
				if editChoice == 1 {
					fmt.Println("Type task title: ")
					title, _ := reader.ReadString('\n')
					title = strings.TrimSpace(title)
					task.Title = title
					fmt.Println(utils.TaskToJsonString(service.UpdateTask(*task)))
				} else {
					var isTaskFinished string
					fmt.Print("Task was finished Y/N? ")
					fmt.Scanln(&isTaskFinished)
					if strings.ToUpper(isTaskFinished) == "Y" {
						task.Status = true
					} else {
						task.Status = false
					}
					fmt.Println(utils.TaskToJsonString(service.UpdateTask(*task)))
				}
			} else {
				fmt.Println("Type task title: ")
				title, _ := reader.ReadString('\n')
				title = strings.TrimSpace(title)
				task.Title = title
				var isTaskFinished string
				fmt.Print("Task was finished Y/N? ")
				fmt.Scanln(&isTaskFinished)
				if strings.ToUpper(isTaskFinished) == "Y" {
					task.Status = true
				} else {
					task.Status = false
				}
			}
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
