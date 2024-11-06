/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"log"

	"github.com/flow-cli/internal/kubernetes"
	"github.com/spf13/cobra"
)

var inputNamespace string

var KubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Interact with your kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func getTargetNamespace() string {
	_, currentNamespace, err := kubernetes.GetCurrentContexts()

	if err != nil {
		log.Fatalf("could not run kubernetes command. Something went wrong with connecting to a cluster. %s", err.Error())
	}

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
