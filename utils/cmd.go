package utils

import (
	"github.com/spf13/cobra"
	"errors"
	"strconv"
	"fmt"
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
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewUpdateCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewMarkStatusDoneCmd())
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

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long: `List all tasks. You can filter tasks by status
    			Example:
    			task_tracker list todo
    			task_tracker list in_progress
    			task_tracker list done
				task_tracker list blocked
    	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) error {
	if len(args) > 0 {
		status := TaskStatus(args[0])
		return ListTasks(status)
	}

	return ListTasks("all")
}

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task",
		Long: `Update a task by providing the task ID and the new status
    Example:
    task_tracker update 1 'new description'
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}

	return cmd
}

func RunUpdateTaskCmd(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Please provide a task id and the new description")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	newDescription := args[1]
	return UpdateTaskDescription(taskIDInt, newDescription)
}

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Long: `Delete a task by providing the task ID

    		Example:
    		task_tracker delete 1
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}

	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a task ID")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	return DeleteTask(taskIDInt)
}

func NewMarkStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, StatusDone)
		},
	}
	return cmd
}

func RunUpdateStatusCmd(args []string, status TaskStatus) error {
	if len(args) == 0 {
		return fmt.Errorf("task ID is required")
	}

	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return UpdateTaskStatus(id, status)
}