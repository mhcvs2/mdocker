package utils

import (
	"bytes"
	"strings"
	"os/exec"
	"strconv"
)

func Exec_shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func OutIsNil(s string) bool {
	cmd := s + "|wc -l"
	out, err := Exec_shell(cmd)
	if err != nil {
		return false
	}
	if outInt, _ := strconv.Atoi(out); outInt == 0 {
		return false
	}
	return true
}

