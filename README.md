# Task Tracker CLI

A simple command-line tool to track and manage your tasks, built in Go.

Project page: https://roadmap.sh/projects/task-tracker

note: Started learning GO a month ago and wanted to start working on projects, used claude in a few places when i got stuck to understand the issue but none of the code is copy pasted or AI generated.

## Features

- Add, update, and delete tasks
- Mark a task as in-progress or done
- List all tasks
- List tasks filtered by status (`done`, `todo`, `in-progress`)
- Tasks persist between runs in a local `storage/tasks.json` file

## Requirements

- Go 1.21 or later installed ([go.dev/dl](https://go.dev/dl))

## Setup

Clone the repository and move into the project folder:

```bash
git clone https://github.com/RishabhWDB/Task-Tracker-v1.git
cd Task-Tracker-v1
```

## Build

```bash
go build -o task-cli
```

On Windows, build with the `.exe` extension:

```powershell
go build -o task-cli.exe
```

This produces a standalone executable in the project folder. A `tasks.json` file will be created automatically inside `storage/` the first time you add a task.

## Usage

```bash
# Adding a new task
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./task-cli update 1 "Buy groceries and cook dinner"
./task-cli delete 1

# Marking a task as in progress or done
./task-cli mark-in-progress 1
./task-cli mark-done 1

# Listing all tasks
./task-cli list

# Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

On Windows, replace `./task-cli` with `./task-cli.exe` in the commands above.

## Task Properties

Each task stored in `storage/tasks.json` has the following properties:

- `Id` — a unique identifier for the task
- `Description` — a short description of the task
- `Status` — the status of the task (`todo`, `in-progress`, `done`)
- `CreatedAt` — the date and time the task was created
- `UpdatedAt` — the date and time the task was last updated

## Project Structure

```
Task-Tracker-v1/
├── main.go            # Entry point — wires everything together
├── commands/           # Task CRUD logic
│   ├── add.go
│   ├── create_id.go
│   ├── delete.go
│   ├── filter_tasks.go
│   ├── mark_done.go
│   ├── mark_in_progress.go
│   └── update_task.go
├── readWrite/          # Command-line argument parsing and JSON persistence
│   ├── jsonStuff.go
│   └── read_input.go
├── types/               # Task and Tasks type definitions, display formatting
│   └── types.go
└── storage/
    └── tasks.json       # Created automatically on first use
```
