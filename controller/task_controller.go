package controller

import (
	"bufio"
	"fmt"
	"go-console-app/entity"
	"go-console-app/service"
	"go-console-app/utils"
	"os"
	"strings"
)

type TaskController struct {
	reader *bufio.Reader
}

func NewTaskController() *TaskController {
	return &TaskController{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (tc *TaskController) GetTaskById() {
	var id string
	fmt.Print("Type task Id: ")
	fmt.Scanln(&id)
	response := service.GetTaskById(id)
	fmt.Println(utils.TaskToJsonString(*response))
}

func (tc *TaskController) GetAllTasks() {
	fmt.Println("All available tasks:")
	response := service.GetAllTasks()
	fmt.Println(response)
}

func (tc *TaskController) CreateTask() {
	fmt.Print("Type task title: ")
	title, _ := tc.reader.ReadString('\n')
	title = strings.TrimSpace(title)
	response := service.CreateTask(title)
	fmt.Println(response)
}

func (tc *TaskController) UpdateTask() {
	var id string
	fmt.Print("Type task Id: ")
	fmt.Scanln(&id)
	task := service.GetTaskById(id)

	var answer string
	fmt.Print("Wanna edit both status and title? (Y/N): ")
	fmt.Scanln(&answer)

	if strings.ToUpper(answer) != "Y" {
		fmt.Print("Type 1 to update task title and 2 to update task status: ")
		var editChoice int
		fmt.Scanln(&editChoice)
		if editChoice == 1 {
			tc.updateTitle(task)
		} else {
			tc.updateStatus(task)
		}
	} else {
		tc.updateTitle(task)
		tc.updateStatus(task)
	}
}

func (tc *TaskController) updateTitle(task *entity.Task) {
	fmt.Print("Type task title: ")
	title, _ := tc.reader.ReadString('\n')
	title = strings.TrimSpace(title)
	task.Title = title
	fmt.Println(utils.TaskToJsonString(service.UpdateTask(*task)))
}

func (tc *TaskController) updateStatus(task *entity.Task) {
	var isTaskFinished string
	fmt.Print("Task was finished (Y/N)? ")
	fmt.Scanln(&isTaskFinished)
	task.Status = strings.ToUpper(isTaskFinished) == "Y"
	fmt.Println(utils.TaskToJsonString(service.UpdateTask(*task)))
}

func (tc *TaskController) DeleteTask() {
	fmt.Print("Type task id: ")
	var id string
	fmt.Scanln(&id)
	fmt.Println(utils.TaskToJsonString(*service.DeleteTaskById(id)))
}
