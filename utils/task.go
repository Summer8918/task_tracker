package utils

import (
    // "fmt"
    "time"
    // "encoding/json"
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
