/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"flow/cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func getImages(namespace string) {
	kubectl := getKubectlPath()
	cmd := fmt.Sprintf(`%s get deployments -o=custom-columns=NAME:.metadata.name,IMAGE:.spec.template.spec.containers[0].image --no-headers=true`, kubectl)

	if namespace != "" {
		cmd = fmt.Sprintf("%s -n %s", cmd, namespace)
	}

	utils.RunCommand(cmd, true)
}

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Fetches a overview of names and images for namespace",
	Run: func(cmd *cobra.Command, args []string) {
		getImages(namespace)
	},
}
