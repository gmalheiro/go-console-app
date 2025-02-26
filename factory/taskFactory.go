package factory

import "go-console-app/entity"

func CreateTask(title string, id int, status bool) entity.Task {
	return entity.Task{
		ID:     id,
		Title:  title,
		Status: status,
	}
}
