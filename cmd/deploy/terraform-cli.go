package deploy

import (
	"flow/cli/utils"
	"fmt"
	"strings"
)

func handleWorkspace(environment string) {
	command := "terraform workspace show"
	output, _ := utils.RunCommandWithOutput(command, false)

	workspace := strings.TrimSpace(output)
	if workspace != environment {
		command = fmt.Sprintf("terraform workspace select %s", environment)
		utils.RunCommand(command)
	}
}

func tfPlan(environment string) {
	handleWorkspace(environment)
	command := fmt.Sprintf("terraform plan --var-file=./vars/%s.tfvars", environment)

	utils.RunCommand(command)
}

func tfApply(environment string, approve bool) {
	handleWorkspace(environment)
	command := fmt.Sprintf("terraform apply --var-file=./vars/%s.tfvars", environment)
	if approve {
		command = command + " -auto-approve"
	}

	utils.RunCommand(command)
}
