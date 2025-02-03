package service

import (
	"fmt"

	"go-console-app/repository"
)

func createTask() string {
	return ""
}

func GetTaskById(id int) string {
	task := repository.GetTaskById(id)

	if task == nil {
		return "Task not found"
	}

	return fmt.Sprintf("ID: %d, Título: %s, Status: %t", task.ID, task.Title, task.Status)
}

func deleteTaskById() string {
	return ""
}

func updateTaskById() string {
	return ""
}

func getAllTasks() string {
	return ""
}
