// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wsl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

func IsWSL() bool {
	var isWSL bool
	isWSL = false
	if runtime.GOOS == "linux" {
		content, err := ioutil.ReadFile("/proc/version")
		if err == nil {
			s := string(content)
			if strings.Contains(s, "WSL2") {
				isWSL = true
			}
		}
	}
	return isWSL
}

func GetWSlTemp() (string, error) {

	var stdout, stderr bytes.Buffer
	args := make([]string, 3)
	args[0] = "/c"
	args[1] = "echo"

	args[2] = "%TEMP%"
	command := exec.Command("cmd.exe", args...)
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()

	stderrString := strings.TrimSpace(stderr.String())
	stdoutString := strings.TrimSpace(stdout.String())

	if _, ok := err.(*exec.ExitError); ok {
		err = fmt.Errorf("Error getting wsl TEMP dir (exec): %s", stderrString)
		return "", err
	}

	if len(stdoutString) == 0 && len(stderrString) > 0 {
		err = fmt.Errorf("Error getting wsl TEMP dir (stderr): \"%s\", %s", stdoutString, stderrString)
		return "", err
	}

	return stdoutString, err
}

func ConvertWindowsPathToWSlPath(winPath string) (string, error) {

	var stdout, stderr bytes.Buffer
	args := make([]string, 3)
	args[0] = "-a"
	args[1] = "-u"

	args[2] = winPath
	command := exec.Command("wslpath", args...)
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()

	stderrString := strings.TrimSpace(stderr.String())

	if _, ok := err.(*exec.ExitError); ok {
		err = fmt.Errorf("wslpath error: %s", stderrString)
		return "", err
	}

	if len(stderrString) > 0 {
		err = fmt.Errorf("wslpath error: %s", stderrString)
		return "", err
	}

	return strings.TrimSpace(stdout.String()), err
}

func ConvertWSlPathToWindowsPath(wslPath string) (string, error) {

	var stdout, stderr bytes.Buffer
	args := make([]string, 3)
	args[0] = "-a"
	args[1] = "-w"

	args[2] = wslPath
	command := exec.Command("wslpath", args...)
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()

	stderrString := strings.TrimSpace(stderr.String())

	if _, ok := err.(*exec.ExitError); ok {
		err = fmt.Errorf("wslpath error: %s", stderrString)
		return "", err
	}

	if len(stderrString) > 0 {
		err = fmt.Errorf("wslpath error: %s", stderrString)
		return "", err
	}

	return strings.TrimSpace(stdout.String()), err
}
