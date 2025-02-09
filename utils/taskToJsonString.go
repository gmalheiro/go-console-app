package utils

import (
	"encoding/json"
	"fmt"

	"go-console-app/model"
)

func TaskToJsonString(task model.Task) string {
	jsonTask, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Error while converting to JSON: ", err)
		return `{"error": "Internal server error"}`
	}

	return string(jsonTask)
}
