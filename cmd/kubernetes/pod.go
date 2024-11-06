package kubernetes

import (
	"github.com/flow-cli/internal/kubernetes"
	"github.com/spf13/cobra"
)

var podCmd = &cobra.Command{
	Use:     "pods",
	Aliases: []string{"pod", "p"},
	Short:   "Pod Overview",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var podListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List available pods",
	Run: func(cmd *cobra.Command, args []string) {
		kubernetes.ListPods()
	},
}

var deployPodCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy pod to cluster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		namespace := getTargetNamespace()

		kubernetes.DeployPod(input, namespace)
	},
}

var deleteFlowPodsCmd = &cobra.Command{
	Use:     "delete-all",
	Aliases: []string{"del-all"},
	Short:   "Delete all pods created by this cli",
	Run: func(cmd *cobra.Command, args []string) {
		kubernetes.DeleteFlowPods()
	},
}

func init() {
	podCmd.AddCommand(podListCmd)
	podCmd.AddCommand(deployPodCmd)
	podCmd.AddCommand(deleteFlowPodsCmd)
}
