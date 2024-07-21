/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"github.com/flow-cli/internal/kubernetes"
	"github.com/spf13/cobra"
)

var inputNamespace string
var _, currentNamespace string = kubernetes.GetCurrentContexts()

var KubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Interact with your kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func getTargetNamespace() string {
	if inputNamespace == "" {
		return currentNamespace
	}
	return inputNamespace
}

func init() {
	KubernetesCmd.Aliases = []string{"k", "kube"}
	KubernetesCmd.AddCommand(imagesCmd)
	KubernetesCmd.AddCommand(watchCmd)
	KubernetesCmd.AddCommand(podCmd)

	KubernetesCmd.PersistentFlags().StringVarP(&inputNamespace, "namespace", "n", "", "Select the namespace in your cluster")
}
