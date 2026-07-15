package commands

import (
	"Task-Tracker-v1/types"
	"fmt"
)

func FilterTasks(tasks types.Tasks, status string) {
	if status == "" {
		for _, t := range tasks {
			fmt.Println(t)
		}
	} else {
		for _, t := range tasks {
			if t.Status == status {
				fmt.Println(t)
			}
		}
	}
}
