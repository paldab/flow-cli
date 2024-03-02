package task

import (
	"log"

	"github.com/flow-cli/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",

	Run: func(cmd *cobra.Command, args []string) {

	taskName, err := cmd.Flags().GetString("task")
	if err != nil {
		log.Fatal(err.Error())
	}

	taskType, err := cmd.Flags().GetString("type")

	if err != nil {
		log.Fatal(err.Error())
	}

	task.AddTask(taskName, taskType)
	},
}

func init() {
	addCmd.Flags().String("task", "", "Name of the task")
	addCmd.Flags().String("type", "", "Type of task. Feature, Bugfix, etc...")
}
