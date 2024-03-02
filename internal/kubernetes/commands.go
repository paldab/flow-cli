package kubernetes

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/flow-cli/internal/cli"
)

func GetKubectlPath() string {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatal(err.Error())
	}

	return path
}

func GetImages(namespace string) {
	kubectl := GetKubectlPath()
	cmd := fmt.Sprintf(`%s get deployments -o=custom-columns=NAME:.metadata.name,IMAGE:.spec.template.spec.containers[0].image --no-headers=false`, kubectl)

	if namespace != "" {
		cmd = fmt.Sprintf("%s -n %s", cmd, namespace)
	}

	cli.RunCommand(cmd)
}

func RevertDeployment(resource string){
	command := fmt.Sprintf("kubectl rollout undo deployment %s", resource)
	cli.RunCommand(command)
}

