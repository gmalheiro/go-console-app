package service

import (
	"fmt"
	"strconv"

	"go-console-app/factory"
	"go-console-app/model"
	"go-console-app/repository"
)

func createTask(id int, title string, status bool) string {
	task := factory.CreateTask(title, id, status)
	return fmt.Sprintf("{ID: %d, Title: %s, Status: %t}", task.ID, task.Title, task.Status)
}

func GetTaskById(id string) string {
	taskId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error while converting: ", err)
		return ""
	}

	task := repository.GetTaskById(taskId)

	if task == nil {
		return "Task not found"
	}

	return fmt.Sprintf("{ID: %d, Title: %s, Status: %t}", task.ID, task.Title, task.Status)
}

func deleteTaskById() string {
	return ""
}

func updateTaskById() string {
	return ""
}

func GetAllTasks() []string {
	var stringTasks []string
	var tasks []model.Task = repository.GetAllTasks()
	for i := range tasks {
		task := fmt.Sprintf(
			"{ID: %d, Title: %s, Status: %t}",
			tasks[i].ID,
			tasks[i].Title,
			tasks[i].Status,
		)
		stringTasks = append(stringTasks, task)
	}
	return stringTasks
}
