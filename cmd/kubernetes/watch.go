/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"flow/cli/utils"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/spf13/cobra"
)

var previousOutput string

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

	output, _ := utils.RunCommandWithOutput(cmd, false)
	if output != previousOutput {
		fmt.Println(output)
		previousOutput = output
	}
}

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watches resources and refreshes every X seconds",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nodeList := []string{"node", "nodes", "n"}
		podList := []string{"pod", "pods", "p"}

		command := args[0]
		isNodeCommand := slices.Contains(nodeList, command)
		isPodCommand := slices.Contains(podList, command)

		if !isNodeCommand && !isPodCommand {
			log.Fatal("Your command isn't a Node or a Pod. Please choose node or pod")
		}

		ticker := time.NewTicker(5 * time.Second)
		stopCh := make(chan struct{})

		go func() {
			watchResources(isPodCommand, namespace)
			for {
				select {
				case <-ticker.C:
					watchResources(isPodCommand, namespace)
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
