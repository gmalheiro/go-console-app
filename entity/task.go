package entity

import "strconv"

type Task struct {
	ID     int
	Title  string
	Status bool
}

func TaskToStringSlice(task Task) []string {
	return []string{strconv.Itoa(task.ID), task.Title, strconv.FormatBool(task.Status)}
}
