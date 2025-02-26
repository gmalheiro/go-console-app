package utils

import (
	"encoding/json"
	"fmt"

	"go-console-app/entity"
)

func TaskToJsonString(task entity.Task) string {
	jsonTask, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Error while converting to JSON: ", err)
		return `{"error": "Internal server error"}`
	}

	return string(jsonTask)
}
