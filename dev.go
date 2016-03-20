package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func freeDev() (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("losetup", "-f")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
