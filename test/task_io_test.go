package test

import (
	"encoding/json"
	"testing"
	utils "github.com/Summer8918/task_tracker/utils"
	"os"
	"path/filepath"
	"time"
)

/*
Create a temporary directory unique to this test. The testing package will auto-delete it 
after the test finishes (even on failure).
*/
// withTempCwd changes into a temp directory for isolation and returns a restore func.
func withTempCwd(t *testing.T) func() {
	// t.Helper() tells the test framework “this is a helper.” 
	// If this function fails, the error line reported points to the caller’s line in the test, 
	// not inside the helper — much nicer for debugging.
	t.Helper()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	tmp := t.TempDir()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir temp: %v", err)
	}
	return func() { _ = os.Chdir(orig) }
}

func TestTasksFilePath_UsesCwdAndTasksJSON(t *testing.T) {
	restore := withTempCwd(t)
	defer restore()

	got, err := utils.TasksFilePath()
	if err != nil {
		t.Fatalf("TasksFilePath error: %v", err)
	}

	want := filepath.Join(mustGetwd(t), "tasks.json")
	if got != want {
		t.Fatalf("path mismatch: got %q, want %q", got, want)
	}
}

func TestReadTasksFromFile_CreatesFileAndReturnsEmptySliceWhenMissing(t *testing.T) {
	restore := withTempCwd(t)
	defer restore()

	// Ensure file does not exist
	p, _ := utils.TasksFilePath()
	if _, err := os.Stat(p); err == nil {
		t.Fatalf("expected tasks.json to not exist at start")
	}

	tasks, err := utils.ReadTasksFromFile()
	if err != nil {
		t.Fatalf("ReadTasksFromFile error: %v", err)
	}
	if len(tasks) != 0 {
		t.Fatalf("expected empty slice, got %d", len(tasks))
	}

	// File should now exist with "[]"
	data, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("ReadFile(%s): %v", p, err)
	}
	trim := string(data)
	if trim != "[]" && trim != "[]\n" {
		t.Fatalf("expected file content to be [] or []\\n, got %q", trim)
	}
}

func TestWriteThenReadTasks_RoundTrip(t *testing.T) {
	restore := withTempCwd(t)
	defer restore()

	now := time.Now().UTC().Truncate(time.Second)
	input := []utils.Task{
		{
			ID:          1,
			Description: "first task",
			// Status left as zero-value to avoid depending on the concrete type
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:          2,
			Description: "second task",
			CreatedAt:   now.Add(1 * time.Second),
			UpdatedAt:   now.Add(2 * time.Second),
		},
	}

	if err := utils.WriteTasksToFile(input); err != nil {
		t.Fatalf("WriteTasksToFile error: %v", err)
	}

	// Read back
	output, err := utils.ReadTasksFromFile()
	if err != nil {
		t.Fatalf("ReadTasksFromFile error: %v", err)
	}

	if len(output) != len(input) {
		t.Fatalf("len mismatch: got %d, want %d", len(output), len(input))
	}

	// Spot-check a couple of fields
	if output[0].ID != input[0].ID || output[0].Description != input[0].Description {
		t.Fatalf("first task mismatch: got %+v, want %+v", output[0], input[0])
	}
	if !output[0].CreatedAt.Equal(input[0].CreatedAt) || !output[0].UpdatedAt.Equal(input[0].UpdatedAt) {
		t.Fatalf("first task times mismatch: got %+v, want %+v", output[0], input[0])
	}
}

func TestReadTasksFromFile_BadJSONReturnsError(t *testing.T) {
	restore := withTempCwd(t)
	defer restore()

	p, err := utils.TasksFilePath()
	if err != nil {
		t.Fatalf("tasksFilePath error: %v", err)
	}

	// Write invalid JSON
	if err := os.WriteFile(p, []byte("{not json"), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	tasks, err := utils.ReadTasksFromFile()
	if err == nil {
		t.Fatalf("expected error for bad JSON, got nil (tasks=%v)", tasks)
	}
}

func TestWriteTasksToFile_WritesValidJSON(t *testing.T) {
	restore := withTempCwd(t)
	defer restore()

	in := []utils.Task{
		{ID: 42, Description: "json check", CreatedAt: time.Unix(0, 0).UTC(), UpdatedAt: time.Unix(0, 0).UTC()},
	}

	if err := utils.WriteTasksToFile(in); err != nil {
		t.Fatalf("WriteTasksToFile error: %v", err)
	}

	p, _ := utils.TasksFilePath()
	raw, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	var decoded []utils.Task
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal written JSON: %v (raw=%q)", err, string(raw))
	}
	if len(decoded) != 1 || decoded[0].ID != 42 {
		t.Fatalf("decoded content mismatch: %+v", decoded)
	}
}

// --- helpers ---

func mustGetwd(t *testing.T) string {
	t.Helper()
	p, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	return p
}