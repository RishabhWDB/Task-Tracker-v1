package readWrite

import (
	"Task-Tracker-v1/types"
	"encoding/json"
	"os"
)

func ReadJson() (types.Tasks, error) {
	var tasks types.Tasks
	data, err := os.ReadFile("storage/tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return tasks, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func WriteJson(tasks *types.Tasks) error {
	err := os.MkdirAll("storage", 0755)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("storage/tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
