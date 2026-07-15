package commands

import (
	"Task-Tracker-v1/types"
	"time"
)

func AddTask(tasks types.Tasks, description string) types.Tasks {
	return append(tasks, types.Task{Id: CreateId(tasks), Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()})
}
