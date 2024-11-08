package kubernetes

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetControllerObjects(namespace string, includeDaemonSets bool) []KubernetesControllerObjects {
	client, err := newKubeClient()

	if err != nil {
		log.Fatal(err.Error())
	}

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
	fmt.Println("NAME\tIMAGE\tREPLICAS")

	for _, obj := range objects {
		objectEntry := fmt.Sprintf("%s\t%s\t%d", obj.Name, obj.Image, obj.Replicas)
		fmt.Println(objectEntry)
	}
}
