package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func newBuildCommand() *cobra.Command {
	var runCommand = &cobra.Command{
		Use:                "build <image_name>",
		Short:              "build an image",
		Long:               "Build command to build an image",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		DisableFlagParsing: true,
		RunE:               buildAction,
	}

	return runCommand
}

func buildAction(cmd *cobra.Command, args []string) error {
	runArgs := "DOCKER_CONFIG=/local/.docker nerdctl "
	runArgs += cmd.Name()

	for _, arg := range args {
		if _, err := os.Stat(arg); err == nil {
			fullPath, _ := filepath.Abs(arg)
			homeDir, _ := os.UserHomeDir()
			newPath := strings.Replace(fullPath, homeDir, "/mnt", 1)
			arg = newPath
		}
		runArgs += " "
		runArgs += arg
		fmt.Printf("arg: %v \n", arg)
	}
	limaArgs := []string{"shell", "br", "sudo", "sheltie", "bash", "-c", runArgs}

	runCmd := exec.Command("limactl", limaArgs...)
	fmt.Println("cmd: ", runCmd)
	runCmd.Stdin = os.Stdin
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	return runCmd.Run()
}
