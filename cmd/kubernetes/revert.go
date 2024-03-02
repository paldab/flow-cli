package kubernetes

import (
	"github.com/flow-cli/internal/kubernetes"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert",
	Args:  cobra.ExactArgs(1),
	Short: "Reverts a deployment to the previous version",
	Run: func(cmd *cobra.Command, args []string) {
		kubernetes.RevertDeployment(args[0])
	},
}
