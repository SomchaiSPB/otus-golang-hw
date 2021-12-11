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

	cmd1 := cmd[0]
	cmd2 := cmd[1]
	arg1 := cmd[2]
	arg2 := cmd[3]

	c := exec.Command(cmd1, cmd2, arg1, arg2)

	c.Stdout = os.Stdout

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}

	return code
}
