package task

import (
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task keeps track if the tasks that you have done. Use one of it's subfunctions!",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TaskCmd.AddCommand(addCmd)
	TaskCmd.AddCommand(listCmd)
	TaskCmd.AddCommand(createBackupCmd)
}
