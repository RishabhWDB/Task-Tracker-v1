package main

import (
	"Task-Tracker-v1/commands"
	"Task-Tracker-v1/input"
	"Task-Tracker-v1/types"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	var tasks types.Tasks

	data, err := os.ReadFile("storage/tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No tasks.json found.")
		} else {
			fmt.Println(err)
		}
	} else {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println(err)
		}
	}

	cmd, err := input.ReadInput()
	if err != nil {
		fmt.Println(err)
	} else {
		switch cmd {
		case "add":
			if len(args) == 3 {
				tasks = commands.AddTask(tasks, args[2])
				fmt.Println("Task added successfully")
				for _, task := range tasks {
					fmt.Printf("%+v\n", task)
				}
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
					for _, task := range tasks {
						fmt.Printf("%+v\n", task)
					}
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
					for _, task := range tasks {
						fmt.Printf("%+v\n", task)
					}
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
					for _, task := range tasks {
						fmt.Printf("%+v\n", task)
					}
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
					for _, task := range tasks {
						fmt.Printf("%+v\n", task)
					}
				}
			} else {
				fmt.Println("Invalid argument length")
			}
		case "list":
			if len(args) == 2 {
				commands.FilterTasks(tasks, "")
			} else if len(args) == 3 {
				commands.FilterTasks(tasks, args[2])
			} else {
				fmt.Println("Invalid argument length")
			}
		}

		data, err = json.Marshal(&tasks)
		if err != nil {
			fmt.Println(err)
		}
		err = os.WriteFile("tasks.json", data, 0644)
	}
}
