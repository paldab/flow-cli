package kubernetes

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesControllerObjects struct {
	Name     string
	Image    string
	Replicas int
}

func (controller KubernetesControllerObjects) mapToTableRow() table.Row {
	return table.Row{controller.Name, controller.Image, controller.Replicas}
}

func getKubeConfig() string {
	defaultKubeFolder := ".kube"
	defaultKubeConfig := "config"
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find home directory")
	}

	return filepath.Join(userHomeDir, defaultKubeFolder, defaultKubeConfig)
}

func GetCurrentContexts() (string, string, error) {
	config, err := clientcmd.LoadFromFile(getKubeConfig())
	if err != nil {
		return "", "", err
	}

	currentContext := config.CurrentContext
	if currentContext == "" {
		return "", "", fmt.Errorf("No current context found in kubeconfig")
	}

	currentContextConfig := config.Contexts[currentContext]
	if currentContextConfig == nil {
		return "", "", fmt.Errorf("Context %s not found in kubeconfig", currentContext)
	}

	currentNamespace := currentContextConfig.Namespace
	if currentNamespace == "" {
		// Default namespace if not set in kubeconfig
		currentNamespace = "default"
	}

	return currentContext, currentNamespace, nil
}

func newKubeClient() (*kubernetes.Clientset, error) {
	kubeConfigPath := getKubeConfig()
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Errorf("could not connect to kubernetes cluster. %s", err.Error())
	}

	client, err := kubernetes.NewForConfig(kubeConfig)

	if err != nil {
		fmt.Errorf("could not create kubernetes client. %s", err.Error())
	}

	return client, nil
}
