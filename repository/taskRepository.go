package repository

import (
	"go-console-app/model"
	"go-console-app/storage"
)

func createTask(task model.Task) {
	storage.Tasks = append(storage.Tasks, task)
}

func GetTaskById(id int) *model.Task {
	for i := range storage.Tasks {
		if storage.Tasks[i].ID == id {
			return &storage.Tasks[i]
		}
	}
	return nil
}

func deleteTaskById() {

}

func updateTaskById() {

}

func GetAllTasks() []model.Task {
	return storage.Tasks
}
