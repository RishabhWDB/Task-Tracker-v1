package commands

import (
	"Task-Tracker-v1/types"
	"errors"
	"strconv"
	"time"
)

func UpdateTask(tasks types.Tasks, id string, description string) (types.Tasks, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return tasks, errors.New("invalid id type")
	}
	for i, t := range tasks {
		if t.Id == n {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, errors.New("id not Found")
}
