package deploy

import (
	"fmt"
	"strings"

	"github.com/flow-cli/internal/cli"
)

func handleWorkspace(environment string) {
	command := "terraform workspace show"
	output, _ := cli.RunCommandWithOutput(command, false)

	workspace := strings.TrimSpace(output)
	if workspace != environment {
		command = fmt.Sprintf("terraform workspace select %s", environment)
		cli.RunCommand(command)
	}
}

func TfPlan(environment string) {
	handleWorkspace(environment)
	command := fmt.Sprintf("terraform plan --var-file=./vars/%s.tfvars", environment)

	cli.RunCommand(command)
}

func TfApply(environment string, approve bool) {
	handleWorkspace(environment)
	command := fmt.Sprintf("terraform apply --var-file=./vars/%s.tfvars", environment)
	if approve {
		command = command + " -auto-approve"
	}

	cli.RunCommand(command)
}
