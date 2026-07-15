package commands

import (
	"Task-Tracker-v1/types"
)

func FilterTasks(tasks types.Tasks, status string) types.Tasks {
	var filtered types.Tasks
	for _, t := range tasks {
		if status == "" || t.Status == status {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
