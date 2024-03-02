package generate

import (
	"fmt"
	"strings"
)

// configuration
func tfVarContent(service, env string) string {
	var content strings.Builder

	content.WriteString(fmt.Sprintf(`project = "csdm-<Project>-%s"
%s_image = ""
%s_host_url = ""

%s_config = {
  replicas = 1
  memory = {
    request = "500Mi"
    limit = "1000Mi"
  }
  cpu = {
    request = "300m"
  }
}
  `, env, service, service, service))

	return content.String()
}

// content for module
func tfModuleMainContent(service string) string {
	return fmt.Sprintf(`terraform {
  backend "gcs" {
    bucket = "tfstate-apps-mab"
    prefix = "%s"
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
    }
    kubectl = {
      source = "alekc/kubectl"
    }
  }
}

resource "kubernetes_namespace" "ns" {
  metadata {
    name = "%s"
  
    labels = {
    environment = local.environment
    terraform   = true
    }
  }
}
`, service, service)
}

func tfLocalsContent(service string) string {
	var content strings.Builder

	content.WriteString(`locals {
  environment = terraform.workspace
}
	`)

	content.WriteString(fmt.Sprintf(`
locals {
  %s-app-config-path = "./config/%s-app/${local.environment}.application.yaml"
}
`, service, service))

	return content.String()
}

func createModuleVariables(service string) []variable {
	serviceVarName := fmt.Sprintf("%s-app", service)

	return []variable{
		{Name: "project", Type: "string"},
		{Name: serviceVarName + "_host_url", Type: "string"},
		{Name: serviceVarName + "_image", Type: "string"},
		{Name: serviceVarName + "_config", Type: `object({
  replicas = number
  memory = object({
    request = optional(string)
	  limit   = optional(string)
  })
  cpu = object({
    request = optional(string)
	  limit   = optional(string)
  })
})`, Default: `{
  replicas = 1
  cpu      = null
  memory   = null
}`},
	}
}

func tfProviderContent() string {
	return `
data "terraform_remote_state" "infrastructure" {
  workspace = terraform.workspace
  backend   = "gcs"
  config = {
    bucket = "tfstate-mab"
  }
}

provider "google" {
  project = var.project
  region  = "europe-west4"
}

data "google_client_config" "default" {}

provider "kubernetes" {
  host                   = "https://${data.terraform_remote_state.infrastructure.outputs.cluster_auth.endpoint}"
  client_certificate     = data.terraform_remote_state.infrastructure.outputs.cluster_auth.client_certificate
  client_key             = data.terraform_remote_state.infrastructure.outputs.cluster_auth.client_key
  cluster_ca_certificate = data.terraform_remote_state.infrastructure.outputs.cluster_auth.ca_certificate
  token                  = data.google_client_config.default.access_token
}

provider "kubectl" {
  host                   = "https://${data.terraform_remote_state.infrastructure.outputs.cluster_auth.endpoint}"
  client_certificate     = data.terraform_remote_state.infrastructure.outputs.cluster_auth.client_certificate
  client_key             = data.terraform_remote_state.infrastructure.outputs.cluster_auth.client_key
  cluster_ca_certificate = data.terraform_remote_state.infrastructure.outputs.cluster_auth.ca_certificate
  token                  = data.google_client_config.default.access_token
  load_config_file       = false
}
`
}

// Content for app
func tfAppMainContent() string {
	var content strings.Builder

	content.WriteString(`terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
    }
  }
}
`)

	return content.String()
}

func tfAppVariablesContent() []variable {
	return []variable{
		{Name: "namespace", Type: "string"},
		{Name: "image", Type: "string"},
		{Name: "configmap_path", Type: "string"},
		{Name: "replicas", Type: "number", Default: "1"},
		{Name: "environment", Type: "string"},
		{Name: "host_url", Type: "string"},
		{Name: "resource_limits", Type: `object({
  cpu    = optional(string)
  memory = optional(string)
})`},
		{Name: "resource_requests", Type: `object({
  cpu    = optional(string)
  memory = optional(string)
})`},
	}
}

// utils
func generateVariablesStringContent(variables []variable) string {
	var content strings.Builder

	for _, v := range variables {
		content.WriteString(fmt.Sprintf(`variable "%s" {
  type = %s`, v.Name, v.Type))

		if v.Default != "" {
			content.WriteString(fmt.Sprintf("\n  default = %s", v.Default))
		}

		content.WriteString("\n}\n\n")
	}

	return content.String()
}
