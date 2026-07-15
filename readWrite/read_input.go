package readWrite

import (
	"errors"
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
