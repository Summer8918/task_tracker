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
    StatusInpProcess   TaskStatus = "in_progress"
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

func statusColor(status TaskStatus) string {
	switch status {
	case StatusTodo:
		return "#3C3C3C"
	case StatusInpProcess:
		return "202"
	case StatusDone:
		return "#04B575"
    case StatusBlocked:
        return "#565656"
	default:
		return "#3C3C3C"
	}
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

func ListTasks(status TaskStatus) error {
    tasks, err := ReadTasksFromFile()
    if err != nil {
        return err
    }

    if len(tasks) == 0 {
        fmt.Println(
            lipgloss.NewStyle().
                Bold(true).
                Padding(1, 0).
                Foreground(lipgloss.Color("#FFCC66")).
                Render("No tasks found"))
        return nil
    }

    var filteredTasks []Task
    switch status {
    case "all":
        filteredTasks = tasks

    case StatusTodo:
        for _, task := range tasks {
            if task.Status == StatusTodo {
                filteredTasks = append(filteredTasks, task)
            }
        }

    case StatusInpProcess:
        for _, task := range tasks {
            if task.Status == StatusTodo {
                filteredTasks = append(filteredTasks, task)
            }
        }

    case StatusDone:
        for _, task := range tasks {
            if task.Status == StatusDone {
                filteredTasks = append(filteredTasks, task)
            }
        }

    case StatusBlocked:
        for _, task := range tasks {
            if task.Status == StatusBlocked {
                filteredTasks = append(filteredTasks, task)
            }
        }
    }

    fmt.Println()
	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			MarginBottom(1).
			Render(fmt.Sprintf("Tasks (%s)", status)))
	for _, task := range filteredTasks {
		formattedId := lipgloss.NewStyle().
			Bold(true).
			Width(5).
			Render(fmt.Sprintf("ID:%d", task.ID))
		formattedStatus := lipgloss.NewStyle().
			Bold(true).
			Width(12).
			Foreground(lipgloss.Color(statusColor(task.Status))).
			Render(string(task.Status))

		relativeUpdatedTime := task.UpdatedAt.Format("2006-01-02 15:04:05")

		taskStyle := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			Render(fmt.Sprintf("%s %s %s (%s)", formattedId, formattedStatus, task.Description, relativeUpdatedTime))
		fmt.Println(taskStyle)
	}
	fmt.Println()

	return nil
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
    for i := range tasks { // iterate by index to modify in place
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			taskExists = true
			break
		}
	}

	if !taskExists {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask updated successfully: %s\n\n", formattedId)
	return WriteTasksToFile(tasks)
}
