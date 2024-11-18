package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func GetShell() string {
	shell := os.Getenv("SHELL")

	var stdout bytes.Buffer
	cmd := exec.Command(shell, "--version")
	cmd.Stdout = &stdout
	err := cmd.Run()

	if err != nil {
		return "Unknown shell"
	}

	return strings.Trim(stdout.String(), "\n")
}
