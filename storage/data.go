package storage

import "go-console-app/model"

var Tasks = []model.Task{
	{ID: 1, Status: false, Title: "Learn go"},
	{ID: 2, Status: true, Title: "Learn Java"},
	{ID: 3, Status: false, Title: "Read a book"},
}
