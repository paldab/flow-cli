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

func init() {
	podCmd.AddCommand(podListCmd)
	podCmd.AddCommand(deployPodCmd)
}

var podListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available pods",
	Run: func(cmd *cobra.Command, args []string) {
		kubernetes.ListPods()
	},
}

var deployPodCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"pod", "p"},
	Short:   "Deploy pod to cluster",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		namespace := getTargetNamespace()

		kubernetes.DeployPod(input, namespace)
	},
}
