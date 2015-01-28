package cmd

import (
	"os/exec"
	"strings"
)

func Exec(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
