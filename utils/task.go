package utils

import (
    "fmt"
    "time"
    // "encoding/json"
    "github.com/charmbracelet/lipgloss"
)

// Define the type
type TaskStatus string

// Define the only values you allow
const (
    StatusTodo    TaskStatus = "todo"
    StatusDoing   TaskStatus = "in_progress"
    StatusDone    TaskStatus = "done"
    StatusBlocked TaskStatus = "blocked"
)

type Task struct {
    ID          int64      `json:"id"`
    Description string     `json:"description"`
    Status      TaskStatus `json:"status"`
    CreatedAt   time.Time  `json:"createdAt"`
    UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
    return &Task{
        ID:          id,
        Description: description,
        Status:      StatusTodo,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var newTaskId int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.ID + 1
	} else {
		newTaskId = 0
	}

	task := NewTask(newTaskId, description)
	tasks = append(tasks, *task)

    // lipgloss (a terminal styling lib) to pretty-print the task ID
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC00"))

	formattedId := style.Render(fmt.Sprintf("(ID: %d)", task.ID))
	fmt.Printf("\nTask added successfully: %s\n\n", formattedId)
	return WriteTasksToFile(tasks)
}