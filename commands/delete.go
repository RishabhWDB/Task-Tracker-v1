package commands

import (
	"Task-Tracker-v1/types"
	"errors"
	"strconv"
)

func Delete(tasks types.Tasks, id string) (types.Tasks, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return tasks, errors.New("invalid id type")
	}
	for i, t := range tasks {
		if t.Id == n {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, errors.New("id not found")
}
