package utils

import (
	"github.com/spf13/cobra"
	"errors"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use: "task_tracker",
		Short: "task_tracker is a CLI tool for managing tasks",
		Long: `task_tracker is a CLI tool for managing tasks. It allows you to create, list, and delete tasks.
		You can also mark tasks as completed and update their status.
		`,
	}
	cmd.AddCommand(NewAddCmd())
	return cmd;
}

//# Adding a new task, task_tracker add "Buy groceries"
func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		// Run but returns an error.
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}

	description := args[0]
	return AddTask(description)
}

