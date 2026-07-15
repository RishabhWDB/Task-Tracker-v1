package commands

import (
	"Task-Tracker-v1/types"
	"errors"
	"strconv"
	"time"
)

func MarkDone(tasks types.Tasks, id string) (types.Tasks, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	for i, t := range tasks {
		if t.Id == n {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, errors.New("id not Found")
}
