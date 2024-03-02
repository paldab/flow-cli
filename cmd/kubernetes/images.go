/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"github.com/flow-cli/internal/kubernetes"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Fetches a overview of names and images for namespace",
	Run: func(cmd *cobra.Command, args []string) {
		kubernetes.GetImages(namespace)
	},
}
