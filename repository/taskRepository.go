package repository

import (
	"go-console-app/data"
	"go-console-app/entity"
)

func CreateTask(task entity.Task) {
	data.Tasks = append(data.Tasks, task)
}

func GetTaskById(id int) *entity.Task {
	for i := range data.Tasks {
		if data.Tasks[i].ID == id {
			return &data.Tasks[i]
		}
	}
	return nil
}

func DeleteTaskById(id int) *entity.Task {
	task := GetTaskById(id)
	if task != nil {
		for i := range data.Tasks {
			if data.Tasks[i].ID == id {
				data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
				return task
			}
		}
	}
	return nil
}

func UpdateTask(task entity.Task) *entity.Task {
	for i := range data.Tasks {
		if data.Tasks[i].ID == task.ID {
			data.Tasks[i].Status = task.Status
			data.Tasks[i].Title = task.Title
			return &data.Tasks[i]
		}
	}
	return nil
}

func GetAllTasks() []entity.Task {
	return data.Tasks
}

func GetSliceSize() int {
	return len(data.Tasks)
}
