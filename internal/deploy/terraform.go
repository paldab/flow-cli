package deploy

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/flow-cli/internal/cli"
)

func isTerraformModule() bool {
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	files, err := os.ReadDir(currDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tf") {
			return true
		}
	}

	return false
}

func handleWorkspace(environment string) {
	if !isTerraformModule() {
		log.Fatal("you are not in a terraform module")
	}

	command := "terraform workspace show"
	output, _ := cli.RunCommandWithOutput(command, false)

	workspace := strings.TrimSpace(output)
	if workspace != environment {
		command = fmt.Sprintf("terraform workspace select %s", environment)
		cli.RunCommand(command)
	}
}

func getVarFile(varFile, environment string) string {
	if varFile == "" {
		return environment
	}

	return varFile
}

func TfPlan(environment, varFile string) {
	handleWorkspace(environment)
	targetVarFile := getVarFile(varFile, environment)
	command := fmt.Sprintf("terraform plan --var-file=./vars/%s.tfvars", targetVarFile)

	cli.RunCommand(command)
}

func TfApply(environment, varFile string, approve bool) {
	handleWorkspace(environment)
	targetVarFile := getVarFile(varFile, environment)

	command := fmt.Sprintf("terraform apply --var-file=./vars/%s.tfvars", targetVarFile)
	if approve {
		command = command + " -auto-approve"
	}

	cli.RunCommand(command)
}
