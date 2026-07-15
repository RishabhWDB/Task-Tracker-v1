package readWrite

import (
	"Task-Tracker-v1/types"
	"encoding/json"
	"os"
)

func ReadJson(tasks types.Tasks) (types.Tasks, error) {
	data, err := os.ReadFile("storage/tasks.json")
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func WriteJson(tasks *types.Tasks) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("storage/tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
