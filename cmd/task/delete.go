/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"log"

	"github.com/spf13/cobra"
)

func deleteTask(target any) {
	return
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {

		lastFlag, err := cmd.Flags().GetBool("last")
		if err != nil {
			log.Fatal(err.Error())
		}

		if lastFlag {
			// Lookup last item from tasks
			deleteTask("Last item")
			return
		}
	},
}

func init() {
	deleteCmd.Flags().BoolP("last", "l", false, "Deletes last task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
