package main

import (
	"go-console-app/service"
)

func main() {
	response := service.GetTaskById(3)
	println(response)
}
