package main

import (
	"Task-Tracker-v1/commands"
	"Task-Tracker-v1/readWrite"
	"errors"
	"fmt"
	"os"
)

func ReadInput() (string, error) {
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
	tasks, err := readWrite.ReadJson()
	if err != nil {
		fmt.Println(err)
	}

	cmd, err := ReadInput()
	if err != nil {
		fmt.Println(err)
	} else {
		switch cmd {
		case "add":
			if len(args) == 3 {
				tasks = commands.AddTask(tasks, args[2])
				fmt.Println("Task added successfully")
				fmt.Println(tasks)
			} else {
				fmt.Println("Invalid argument length")
			}
		case "update":
			if len(args) == 4 {
				tasks, err = commands.UpdateTask(tasks, args[2], args[3])
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Task updated successfully")
					fmt.Println(tasks)
				}
			} else {
				fmt.Println("Invalid argument length")
			}
		case "delete":
			if len(args) == 3 {
				tasks, err = commands.Delete(tasks, args[2])
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Task deleted successfully!")
					fmt.Println(tasks)
				}
			} else {
				fmt.Println("Invalid argument length")
			}
		case "mark-in-progress":
			if len(args) == 3 {
				tasks, err = commands.MarkInProgress(tasks, args[2])
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Task marked in-progress successfully!")
					fmt.Println(tasks)
				}
			} else {
				fmt.Println("Invalid argument length")
			}
		case "mark-done":
			if len(args) == 3 {
				tasks, err = commands.MarkDone(tasks, args[2])
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Task marked done successfully!")
					fmt.Println(tasks)
				}
			} else {
				fmt.Println("Invalid argument length")
			}
		case "list":
			if len(args) == 2 {
				tasks = commands.FilterTasks(tasks, "")
				fmt.Println(tasks)
			} else if len(args) == 3 {
				tasks = commands.FilterTasks(tasks, args[2])
				fmt.Println(tasks)
			} else {
				fmt.Println("Invalid argument length")
			}
		}

		err = readWrite.WriteJson(&tasks)
		if err != nil {
			fmt.Println(err)
		}
	}
}
