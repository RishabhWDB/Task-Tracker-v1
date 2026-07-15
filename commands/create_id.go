package commands

import "Task-Tracker-v1/types"

func CreateId(tasks types.Tasks) int {
	uid := 1
	for _, t := range tasks {
		if t.Id > uid {
			uid = t.Id + 1
		}
	}
	return uid
}
