/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var namespace string

func getKubectlPath() string {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatal(err.Error())
	}

	return path
}

// kubernetesCmd represents the kubernetes command
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
