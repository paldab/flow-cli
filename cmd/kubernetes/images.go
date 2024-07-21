/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"github.com/flow-cli/internal/kubernetes"
	"github.com/spf13/cobra"
)

var includeDaemonSets, allNamespaces bool

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Fetches a overview of names and images for namespace. Use the -n flag to specify a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		var namespace string

		if allNamespaces {
			namespace = ""
		} else {
			namespace = getTargetNamespace()
		}

		objects := kubernetes.GetControllerObjects(namespace, includeDaemonSets)
		kubernetes.ShowControllerObjects(objects)
	},
}

func init() {
	imagesCmd.Flags().BoolVarP(&includeDaemonSets, "daemonSets", "d", false, "daemonSets true includes images of daemonSets")
	imagesCmd.Flags().BoolVarP(&allNamespaces, "all", "a", false, "All Namespaces")
}
