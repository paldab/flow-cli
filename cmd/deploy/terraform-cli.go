package deploy

import (
	"flow/cli/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func handleWorkspace(environment string) {
	command := "terraform workspace show"
	output, _ := utils.RunCommand(command, false)

	workspace := strings.TrimSpace(output)
	if workspace != environment {
		command = fmt.Sprintf("terraform workspace select %s", environment)
	}
}

func runTerraform(command string) {
	cmd := utils.PrepareCommand(command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatal(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err.Error())
	}
}

func tfPlan(environment string) {
	handleWorkspace(environment)
	command := fmt.Sprintf("terraform plan --var-file=./vars/%s.tfvars", environment)

	runTerraform(command)
}

func tfApply(environment string, approve bool) {
	handleWorkspace(environment)
	command := fmt.Sprintf("terraform apply --var-file=./vars/%s.tfvars", environment)
	if approve {
		command = command + " -auto-approve"
	}

	runTerraform(command)
}
