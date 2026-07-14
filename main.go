package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

func AddTask(tasks Tasks, description string) Tasks {
	return append(tasks, Task{Id: CreateId(tasks), Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()})
}

func MarkInProgress(tasks Tasks, id string) (Tasks, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	for i, t := range tasks {
		if t.Id == n {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, errors.New("id not Found")
}

func MarkDone(tasks Tasks, id string) (Tasks, error) {
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

func UpdateTask(tasks Tasks, id string, description string) (Tasks, error) {
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

func Delete(tasks Tasks, id string) (Tasks, error) {
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
	uid := 1
	for _, t := range tasks {
		if t.Id > uid {
			uid = t.Id + 1
		}
	}
	return uid
}

func ReadInput(tasks Tasks) (string, error) {
	args := os.Args
	var initCommands []string
	initCommands = append(initCommands, "add", "update", "delete", "mark-in-progress", "mark-done", "list")
	if len(args) < 2 {
		return "", errors.New("invalid arguments")
	}
	for _, cmd := range initCommands {
		if cmd == args[1] {
			return cmd, nil
		}
	}
	return "", errors.New("command not found")

}

func main() {
	args := os.Args
	var tasks Tasks
	cmd, err := ReadInput(tasks)
	if err != nil {
		fmt.Println(err)
	} else {
		switch cmd {
		case "add":
			tasks = AddTask(tasks, args[2])
			fmt.Println("Task added successfully")
			fmt.Println(tasks)
		case "update":
			tasks, err = UpdateTask(tasks, args[2], args[3])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Task updated successfully")
			fmt.Println(tasks)
		case "delete":
			tasks, err = Delete(tasks, args[2])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Task deleted successfully!")
			fmt.Println(tasks)
		case "mark-in-progress":
			tasks, err = MarkInProgress(tasks, args[2])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Task marked in-progress successfully!")
			fmt.Println(tasks)
		case "mark-done":
			tasks, err = MarkDone(tasks, args[2])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Task marked done successfully!")
			fmt.Println(tasks)
		case "list":
			if len(args) < 3 {
				FilterTasks(tasks, "")
			} else {
				FilterTasks(tasks, args[2])
			}
		}
	}
}
