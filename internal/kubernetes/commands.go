package kubernetes

import (
	"context"
	"log"

	"github.com/flow-cli/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetControllerObjects(namespace string, includeDaemonSets bool) []KubernetesControllerObjects {
	var controllers []KubernetesControllerObjects
	deployments, err := client.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("error getting Deployments: %v\n", err)
	}

	statefulSets, err := client.AppsV1().StatefulSets(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("error getting StatefulSets: %v\n", err)
	}

	for _, deployment := range deployments.Items {
		if len(deployment.Spec.Template.Spec.Containers) == 0 {
			continue
		}

		controllers = append(controllers, KubernetesControllerObjects{
			Name:     deployment.Name,
			Image:    deployment.Spec.Template.Spec.Containers[0].Image,
			Replicas: int(*deployment.Spec.Replicas),
		})
	}

	for _, statefulSet := range statefulSets.Items {
		if len(statefulSet.Spec.Template.Spec.Containers) == 0 {
			continue
		}

		controllers = append(controllers, KubernetesControllerObjects{
			Name:     statefulSet.Name,
			Image:    statefulSet.Spec.Template.Spec.Containers[0].Image,
			Replicas: int(*statefulSet.Spec.Replicas),
		})
	}

	if includeDaemonSets {
		nodes, err := client.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
		nodeCount := len(nodes.Items)
		if err != nil {
			log.Fatalf("error getting Nodes: %v\n", err)
		}

		daemonSets, err := client.AppsV1().DaemonSets(namespace).List(context.Background(), v1.ListOptions{})
		if err != nil {
			log.Fatalf("error getting Daemonsets: %v\n", err)
		}

		for _, daemonSet := range daemonSets.Items {
			if len(daemonSet.Spec.Template.Spec.Containers) == 0 {
				continue
			}

			controllers = append(controllers, KubernetesControllerObjects{
				Name:     daemonSet.Name,
				Image:    daemonSet.Spec.Template.Spec.Containers[0].Image,
				Replicas: nodeCount,
			})
		}
	}

	return controllers
}

func ShowControllerObjects(objects []KubernetesControllerObjects) {
	header := table.Row{"Name", "Image", "Replicas"}
	rows := []table.Row{}

	for _, obj := range objects {
		rows = append(rows, obj.mapToTableRow())
	}

	utils.PrintTable(header, rows)
}
