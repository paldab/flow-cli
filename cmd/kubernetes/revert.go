/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"flow/cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func getLastRevisions(resource string) {
	command := fmt.Sprintf("kubectl rollout history deployment %s --all-namespaces", resource)
	output, _ := utils.RunCommandWithOutput(command, false)
	fmt.Println(output)
}

// fmt.Sprintf("kubectl rollout history deployment %s --revision=<revision-number>", resource)

// revertCmd represents the revert command
var revertCmd = &cobra.Command{
	Use:   "revert",
	Args:  cobra.ExactArgs(1),
	Short: "Reverts a deployment to the previous image",
	Run: func(cmd *cobra.Command, args []string) {
		getLastRevisions(args[0])
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// revertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// revertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
