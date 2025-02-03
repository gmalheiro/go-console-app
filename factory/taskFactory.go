package factory

import "go-console-app/model"

func CreateTask(title string, id int, status bool) model.Task {
	return model.Task{
		ID:     id,
		Title:  title,
		Status: status,
	}
}
