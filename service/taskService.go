package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go-console-app/factory"
	"go-console-app/model"
	"go-console-app/repository"
)

func CreateTask(title string) string {
	id := repository.GetSliceSize() + 1
	task := factory.CreateTask(title, id, false)
	repository.CreateTask(task)
	return fmt.Sprintf("{ID: %d, Title: %s, Status: %t}", task.ID, task.Title, task.Status)
}

func GetTaskById(id string) *model.Task {
	taskId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error while converting: ", err)
		return nil
	}

	task := repository.GetTaskById(taskId)

	if task == nil {
		fmt.Println("Task not found")
		return nil
	}

	return task
}

func DeleteTaskById(id string) *model.Task {
	taskId := GetTaskById(id)
	return repository.DeleteTaskById(taskId.ID)
}

func UpdateTask(task model.Task) model.Task {
	return *repository.UpdateTask(task)
}

func GetAllTasks() string {
	tasks := repository.GetAllTasks()

	jsonTasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error while converting to JSON: ", err)
		return `{"error": "Internal server error"}`
	}

	return string(jsonTasks)
}
