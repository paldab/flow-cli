/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"fmt"
	"log"
	"os/exec"
	"slices"
	"time"

	"github.com/flow-cli/internal/cli"
	"github.com/spf13/cobra"
)

var previousOutput string

// USES KUBECTL CLI SETUP
func getKubectlPath() string {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatalf("could not get kubectl config. %s", err.Error())
	}

	return path
}

func watchResources(isPod bool, namespace string) {
	kubectl := getKubectlPath()
	resource := "node"

	if isPod {
		resource = "pod"
	}

	cmd := fmt.Sprintf("%s top %s", kubectl, resource)

	if namespace != "" {
		cmd = fmt.Sprintf("%s -n %s", cmd, namespace)
	}

	output, _ := cli.RunCommandWithOutput(cmd, false)
	if output != previousOutput {
		fmt.Println(output)
		previousOutput = output
	}
}

var refreshInterval int = 5
var watchCmd = &cobra.Command{
	Use:     "watch",
	Aliases: []string{"top"},
	Short:   fmt.Sprintf("Watches resources and refreshes every %d seconds", refreshInterval),
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nodeList := []string{"node", "nodes", "n"}
		podList := []string{"pod", "pods", "p"}

		command := args[0]
		isNodeCommand := slices.Contains(nodeList, command)
		isPodCommand := slices.Contains(podList, command)

		if !isNodeCommand && !isPodCommand {
			log.Fatal("your command isn't a Node or a Pod.")
		}

		ticker := time.NewTicker(time.Duration(refreshInterval) * time.Second)
		stopCh := make(chan struct{})

		go func() {
			watchResources(isPodCommand, inputNamespace)
			for {
				select {
				case <-ticker.C:
					watchResources(isPodCommand, inputNamespace)
				case stopCh <- struct{}{}:
					ticker.Stop()
					return
				}
			}
		}()
		time.Sleep(time.Minute)
		<-stopCh
	},
}
