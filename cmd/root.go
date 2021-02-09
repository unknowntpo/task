package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a cli task manager",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
