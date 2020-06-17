package cmd

import (
	"fmt"

	"github.com/desferreira/go-challenges/task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		allTasks, err := db.AllTasks()
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range allTasks {
			fmt.Printf("%v\t%v\n", v.Key, v.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
