package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// default Version for local builds
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Flow version: %s\n", Version)
	},
}
