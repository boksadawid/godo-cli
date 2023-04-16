package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "GODO-CLI is a command-line application for managing and scheduling tasks.",
	Long: `GODO-CLI allows you to create one-time, recurring, and backlog tasks,
	and it notifies you when a task is due. If you don't complete a task by its deadline,
	the application blocks certain websites on your computer.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	InitCommands(rootCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
