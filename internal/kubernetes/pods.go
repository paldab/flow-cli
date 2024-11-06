package kubernetes

import (
	"context"
	"fmt"
	"log"
	"os"
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
	fmt.Println("NAME\tIMAGE")
	for _, pod := range availablePods {
		podEntry := fmt.Sprintf("%s\t%s", pod.Name, pod.Image)
		fmt.Println(podEntry)
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

func DeleteFlowPods() {
	client, err := newKubeClient()

	if err != nil {
		log.Fatal(err.Error())
	}

	labelSelector := "managed-by=flow"
	pods, err := client.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{
		LabelSelector: labelSelector,
	})

	if err != nil {
		log.Fatalf("error listing pods: %s", err.Error())
	}

	if len(pods.Items) == 0 {
		fmt.Println("all pods managed by flow are deleted")
		return
	}

	deletePolicy := v1.DeletePropagationForeground // Deletes all dependencies as well
	for _, pod := range pods.Items {
		err := client.CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, v1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting pod: %v\n", err)
		}
	}

	fmt.Println("deleting all pods managed by flow...")
}
