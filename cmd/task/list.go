/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"github.com/flow-cli/internal/task"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks registered",
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTasks()
	},
}
