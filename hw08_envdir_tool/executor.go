package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, value := range env {
		if value.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				log.Println(err)
			}
		}
		err := os.Setenv(key, value.Value)
		if err != nil {
			log.Println(err)
		}
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

	return c.ProcessState.ExitCode()
}
