package generate

import (
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a structure",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	GenerateCmd.Aliases = []string{"g", "gen"}
	GenerateCmd.AddCommand(terraformServiceCmd)
}
