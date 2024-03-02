package generate

import (
	"github.com/flow-cli/internal/generate"
	"github.com/spf13/cobra"
)

var terraformServiceCmd = &cobra.Command{
	Use:   "terraform-service",
	Args:  cobra.ExactArgs(1),
	Short: "Generates a terraform service/application deployed with terraform",
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		generate.GenerateStructure(serviceName)
	},
}

func init() {
	terraformServiceCmd.Aliases = []string{"tfs", "tfservice", "tfapp", "application", "tf", "service"}
}
