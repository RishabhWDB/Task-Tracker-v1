package main

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks []Task

func AddTask(tasks Tasks, task string) Tasks {
	return append(tasks, Task{Id: CreateId(tasks), Description: task, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()})
}

func MarkInProgress(tasks Tasks, id int) (Tasks, error) {
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, errors.New("id not Found")
}

func MarkDone(tasks Tasks, id int) (Tasks, error) {
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, errors.New("id not Found")
}

func UpdateTask(tasks Tasks, id int, description string) (Tasks, error) {
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, errors.New("id not Found")
}

func Delete(tasks Tasks, id int) (Tasks, error) {
	for i, t := range tasks {
		if t.Id == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, errors.New("id not found")
}

func FilterTasks(tasks Tasks, status string) {
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

func CreateId(tasks Tasks) int {
	uid := 0
	for _, t := range tasks {
		if t.Id > uid {
			uid = t.Id + 1
		}
	}
	return uid
}

func main() {

}
