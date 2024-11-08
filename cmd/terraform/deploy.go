/*
Copyright © 2024 Paul <EMAIL ADDRESS>
*/
package terraform

import (
	"github.com/flow-cli/internal/terraform"
	"github.com/spf13/cobra"
)

var plan, approve bool
var varFile string

var TerraformCmd = &cobra.Command{
	Use:     "terraform",
	Aliases: []string{"tf"},
	Args:    cobra.ExactArgs(1),
	Short:   "Deploys the terraform module in the given environment",
	Run: func(cmd *cobra.Command, args []string) {
		environment := args[0]

		if plan {
			terraform.TfPlan(environment, varFile)
			return
		}
		terraform.TfApply(environment, varFile, approve)
	},
}

func init() {
	TerraformCmd.Flags().StringVarP(&varFile, "varFile", "f", "", "terraform varFile to override default val $input.tfvars)")
	TerraformCmd.Flags().BoolVarP(&plan, "plan", "p", false, "Plans terraform terraformment")
	TerraformCmd.Flags().BoolVarP(&approve, "approve", "a", false, "Auto approves terrafrom terraformment")
}
