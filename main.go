package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
    "github.com/Summer8918/task_tracker/utils"
)

func main() {
    rootCmd := utils.NewRootCmd()
    fmt.Println("Hello task_tracker")
	if err := rootCmd.Execute(); err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Padding(1, 2).
			Bold(true).
			Render(fmt.Sprintf("Error: %s", err))
		fmt.Println(errorStyle)
	}
}
