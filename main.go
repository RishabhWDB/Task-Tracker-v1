package main

import (
	"errors"
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

func AddTask(tasks Tasks, id int, task string) Tasks {
	return append(tasks, Task{Id: id, Description: task, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()})
}

func UpdateStatus(tasks Tasks, id int, status string) (Tasks, error) {
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = status
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
	taskId := 1
	taskId += 1
	//var tasks Tasks
	//switch
}
