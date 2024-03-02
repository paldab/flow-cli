/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"github.com/spf13/cobra"
)

var namespace string

var KubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Interact with your kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	KubernetesCmd.Aliases = []string{"k", "kube"}
	KubernetesCmd.AddCommand(imagesCmd)
	KubernetesCmd.AddCommand(watchCmd)
	KubernetesCmd.AddCommand(revertCmd)

	KubernetesCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Select the namespace in your cluster")
}
