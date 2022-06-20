/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {
	if err := newApp().Execute(); err != nil {
		handleExitCoder(err)
	}
}

type ExitCoder interface {
	error
	ExitCode() int
}

func handleExitCoder(err error) {
	if err == nil {
		return
	}

	if exitErr, ok := err.(ExitCoder); ok {
		os.Exit(exitErr.ExitCode())
		return
	}
}

func newApp() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "washington",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	rootCmd.AddCommand(newBuildCommand(),
		createCommand("run"),
		createCommand("images"),
		createCommand("image"),
		createCommand("rmi"),
		createCommand("ps"),
		createCommand("rm"),
		createCommand("pull"),
		createCommand("stop"),
		createCommand("exec"),
		createCommand("create"),
		createCommand("start"))
	return rootCmd
}
