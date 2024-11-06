package kubernetes

import (
	"context"
	"fmt"
	"log"
	"strings"

	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodData struct {
	Name    string
	Image   string
	Command []string
}

var availablePods []PodData = []PodData{
	{Name: "ping", Image: "busybox", Command: []string{"sleep", "3600"}},
	{Name: "curl", Image: "curlimages/curl", Command: []string{"sleep", "3600"}},
	{Name: "gcloud", Image: "google/cloud-sdk:slim", Command: []string{"sleep", "3600"}},
}

func ListPods() {
	for _, pod := range availablePods {
		fmt.Println(pod.Name)
	}
}

func SearchPod(input string) PodData {
	for _, pod := range availablePods {
		if strings.Contains(pod.Name, input) {
			return pod
		}
	}

	log.Fatal("Could not find any predefined pods with that name!")
	return PodData{}
}

func createPod(pod PodData, namespace string) *core.Pod {
	return &core.Pod{
		ObjectMeta: v1.ObjectMeta{
			Name:      pod.Name,
			Namespace: namespace,
			Labels: map[string]string{
				"managed-by": "flow",
			},
		},
		Spec: core.PodSpec{
			Containers: []core.Container{
				{
					Name:            pod.Name,
					Image:           pod.Image,
					Command:         pod.Command,
					ImagePullPolicy: "IfNotPresent",
					Ports: []core.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}
}

func DeployPod(input, namespace string) {
    client, err := newKubeClient()

    if err != nil {
        log.Fatal(err.Error())
    }

	targetPod := SearchPod(input)
	newPod := createPod(targetPod, namespace)

	_, err = client.CoreV1().Pods(namespace).Create(context.Background(), newPod, v1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating pod: %v", err)
	}

	fmt.Printf("Pod %s created successfully in namespace %s\n", targetPod.Name, namespace)
}
