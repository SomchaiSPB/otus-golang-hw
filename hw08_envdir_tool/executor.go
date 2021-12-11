package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var code int

	for key, value := range env {
		if value.NeedRemove {
			_ = os.Unsetenv(key)
		}
		_ = os.Setenv(key, value.Value)
	}

	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL)
	//signal.Ignore(syscall.SIGTERM)

	c := exec.Command(cmd[0], cmd[1], cmd[2], cmd[3])

	c.Stdout = os.Stdout

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}

	return code
}
