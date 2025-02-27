package factory

import (
	"go-console-app/entity"
	"strconv"
)

func CreateTask(title string, id int, status bool) entity.Task {
	return entity.Task{
		ID:     id,
		Title:  title,
		Status: status,
	}
}

func CreateStringTask(title string, id string, status string) entity.Task {
	taskId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	taskStatus, err := strconv.ParseBool(status)

	if err != nil {
		panic(err)
	}

	return entity.Task{
		ID:     taskId,
		Title:  title,
		Status: taskStatus,
	}
}
