package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Simple todo manager",
	Long:  "You can insert, delete and list all your tasks for the day inside your terminal! Awesome!",
}
