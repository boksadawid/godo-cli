package cmd

import "github.com/spf13/cobra"

func InitCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(listAllCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(DoneCmd)

}
