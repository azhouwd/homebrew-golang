package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func createCommand(cmdName string) *cobra.Command {
	var command = &cobra.Command{
		Use:                cmdName,
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		DisableFlagParsing: true,
		RunE:               createAction,
	}
	return command
}

func createAction(cmd *cobra.Command, args []string) error {
	nerdctlArgs := "DOCKER_CONFIG=/local/.docker nerdctl "
	nerdctlArgs += cmd.Name()
	for _, arg := range args {
		nerdctlArgs += " "
		nerdctlArgs += arg
	}
	limaArgs := []string{"shell", "br", "sudo", "sheltie", "bash", "-c", nerdctlArgs}

	newCmd := exec.Command("limactl", limaArgs...)
	fmt.Println("cmd: ", newCmd)
	newCmd.Stdin = os.Stdin
	newCmd.Stdout = os.Stdout
	newCmd.Stderr = os.Stderr
	return newCmd.Run()
}
