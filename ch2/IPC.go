package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")

	cmd.Env = []string{"PS1=-[namespace-process]-#"}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		os.Exit(1)
	}
}
