package repository

import (
	"go-console-app/model"
	"go-console-app/storage"
)

func CreateTask(task model.Task) {
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

func UpdateTask(task model.Task) *model.Task {
	for i := range storage.Tasks {
		if storage.Tasks[i].ID == task.ID {
			storage.Tasks[i].Status = task.Status
			storage.Tasks[i].Title = task.Title
			return &storage.Tasks[i]
		}
	}
	return nil
}

func GetAllTasks() []model.Task {
	return storage.Tasks
}

func GetSliceSize() int {
	return len(storage.Tasks)
}
