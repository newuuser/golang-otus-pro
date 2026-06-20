package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	exe := exec.Command(cmd[0], cmd[1:]...) //nolint

	osEnv := os.Environ()
	runEnv := make([]string, 0)

	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	exe.Stdin = os.Stdin

	unset := make(map[string]struct{})

	for k, v := range env {
		if v.NeedRemove {
			unset[k] = struct{}{}
		}
	}

	for _, s := range osEnv {
		split := strings.SplitN(s, "=", 2)
		if _, ok := unset[split[0]]; !ok {
			runEnv = append(runEnv, s)
		}
	}

	for k, v := range env {
		if !v.NeedRemove {
			runEnv = append(runEnv, fmt.Sprintf("%s=%s", k, v.Value))
		}
	}

	exe.Env = runEnv
	err := exe.Start()
	if err != nil {
		log.Print(err)
		return 3
	}
	err = exe.Wait()

	if err == nil {
		return 0
	}

	var e *exec.ExitError
	ok := errors.As(err, &e)
	if ok {
		return e.ExitCode()
	}
	// iff unexpected not ExitError exception like I/O exception
	log.Print(err)
	return 4
}
