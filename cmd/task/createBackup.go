package task

import (
	"github.com/flow-cli/internal/task"
	"github.com/spf13/cobra"
)

var createBackupCmd = &cobra.Command{
	Use:   "createBackup",
	Short: "Creates a backup for the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task.CreateBackup()
	},
}
