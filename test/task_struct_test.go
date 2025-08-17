package test

import (
	"encoding/json"
	"testing"
	"fmt"
	utils "github.com/Summer8918/task_tracker/utils"
)

func TestTask_JSONRoundTrip(t *testing.T) {
	task0 := utils.NewTask(1, "first task");
    // Serialize
    b, err := json.MarshalIndent(task0, "", "  ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))

    // Deserialize back into a struct
    var task1 utils.Task
    if err := json.Unmarshal(b, &task1); err != nil {
    panic(err)
    }
    fmt.Printf("Round-trip OK? %v\n", task1.Description == task0.Description)
}
