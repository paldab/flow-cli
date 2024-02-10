package utils

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func PrepareCommand(command string) *exec.Cmd {
	args := strings.Split(command, " ")
	exe, err := exec.LookPath(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	return exec.Command(exe, args[1:]...)
}

func RunCommand(command string, showOutput bool) (string, string) {
	cmd := PrepareCommand(command)
	var outb, errb bytes.Buffer

	cmd.Stdin = os.Stdin
	cmd.Stderr = &errb
	cmd.Stdout = &outb

	cmd.Run()

	outputs := outb.String()
	errs := errb.String()
	if errs != "" {
		os.Stderr.WriteString(errs)
		return "", errs
	}

	if showOutput {
		os.Stdout.WriteString(outputs)
	}

	return outputs, errs
}
