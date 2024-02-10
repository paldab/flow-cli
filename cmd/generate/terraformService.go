/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package generate

import (
	"flow/cli/utils"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

type structureItem struct {
	Name     string
	Content  string
	Path     string
	IsFolder bool
}

type variable struct {
	Name    string
	Type    string
	Default string
}

func generateVarFiles(service, path string) []structureItem {
	var vars []structureItem
	environments := []string{"dtest", "qa", "staging", "production"}

	for _, env := range environments {
		vars = append(vars, structureItem{
			Name:     fmt.Sprintf("%s.tfvars", env),
			Content:  tfVarContent(service, env),
			IsFolder: false,
			Path:     path,
		})
	}

	return vars
}

func generateConfigFiles(path string) []structureItem {
	var vars []structureItem
	environments := []string{"dtest", "qa", "staging", "production"}

	for _, env := range environments {
		vars = append(vars, structureItem{
			Name:     fmt.Sprintf("%s.application.yaml", env),
			IsFolder: false,
			Path:     path,
		})
	}

	return vars
}

func generateApp(name, path string) []structureItem {
	variablesContent := tfAppVariablesContent()
	varStringContent := generateVariablesStringContent(variablesContent)

	return []structureItem{
		{Name: "main.tf", Content: tfAppMainContent(), IsFolder: false, Path: path},
		{Name: "variables.tf", Content: varStringContent, IsFolder: false, Path: path},
	}
}

func generateStructure(service string) {
	moduleVars := createModuleVariables(service)
	appFolder := fmt.Sprintf("%s-app", service)
	moduleVariablesContent := generateVariablesStringContent(moduleVars)

	tfVarFiles := generateVarFiles(service, "vars")
	appConfigFiles := generateConfigFiles("config")
	baseApp := generateApp(service, appFolder)
	var structure []structureItem = []structureItem{
		{Name: "vars", IsFolder: true},
		{Name: "config", IsFolder: true},
		{Name: appFolder, IsFolder: true},

		{Name: "main.tf", Content: tfModuleMainContent(service), IsFolder: false},
		{Name: "locals.tf", Content: tfLocalsContent(service), IsFolder: false},
		{Name: "providers.tf", Content: tfProviderContent(), IsFolder: false},
		{Name: "variables.tf", Content: moduleVariablesContent, IsFolder: false},
	}

	structure = append(structure, tfVarFiles...)
	structure = append(structure, appConfigFiles...)
	structure = append(structure, baseApp...)

	// Creating service folder
	if err := os.Mkdir(service, os.ModePerm); err != nil {
		log.Fatal("Error creating service directory! " + err.Error())
	}

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Generating structure
	os.Chdir(path.Join(currDir, service))
	for _, item := range structure {
		if item.IsFolder {
			if err := os.Mkdir(item.Name, os.ModePerm); err != nil {
				log.Fatal("Error creating directory. " + err.Error())
			}
			continue
		}

		f, err := os.Create(path.Join(item.Path, item.Name))
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", item.Name, err)
			return
		}
		defer f.Close()

		_, err = f.WriteString(item.Content)
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", item.Name, err)
			return
		}
	}

	// format terraform code
	utils.RunCommand("terraform fmt -recursive", false)

	fmt.Printf("Files generated for service: %s\n", service)
}

var terraformServiceCmd = &cobra.Command{
	Use:   "terraform-service",
	Args:  cobra.ExactArgs(1),
	Short: "Generates a terraform service/application deployed with terraform",
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		generateStructure(serviceName)
	},
}

func init() {
	terraformServiceCmd.Aliases = []string{"tfs", "tfservice", "tfapp", "application", "tf "}
}
