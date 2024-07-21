package kubernetes

import (
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

var client *kubernetes.Clientset = kubeClient()

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

func GetCurrentContexts() (string, string) {
	config, err := clientcmd.LoadFromFile(getKubeConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	currentContext := config.CurrentContext
	if currentContext == "" {
		log.Fatal("No current context found in kubeconfig")
	}

	currentContextConfig := config.Contexts[currentContext]
	if currentContextConfig == nil {
		log.Fatalf("Context %s not found in kubeconfig", currentContext)
	}

	currentNamespace := currentContextConfig.Namespace
	if currentNamespace == "" {
		// Default namespace if not set in kubeconfig
		currentNamespace = "default"
	}

	return currentContext, currentNamespace
}

func kubeClient() *kubernetes.Clientset {
	kubeConfigPath := getKubeConfig()
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	client, err := kubernetes.NewForConfig(kubeConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}
