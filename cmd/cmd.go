package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/unknowntpo/task/db"
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "task",
		Short: "Task is a cli task manager",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
}

func newAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Adds a task to your task list",
		Run: func(cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			err := db.CreateTask(db.Db, task)
			if err != nil {
				fmt.Errorf("Fail to create task: %v", err)
				return
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Add \"%s\" to your task list.\n", task)
		},
	}
}

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list called")
		},
	}
}
func newDoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "do",
		Short: "Mark a task as complete",
		Run: func(cmd *cobra.Command, args []string) {
			var ids []int
			for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Println("Failed to parse the argument: ", arg)
				} else {
					ids = append(ids, id)
				}
			}
			fmt.Println("ids: ", ids)
		},
	}
}

// Init inits an root command and add all commands in cmd pkg.
// return error if it exist
func Init() (*cobra.Command, error) {
	rootCmd := newRootCmd()
	listCmd := newListCmd()
	addCmd := newAddCmd()
	doCmd := newDoCmd()

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(doCmd)

	rootCmd.Execute()

	return rootCmd, nil
}
