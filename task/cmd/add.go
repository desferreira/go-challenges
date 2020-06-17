package cmd

import (
	"fmt"
	"strings"

	"github.com/desferreira/go-challenges/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks to the list. Simple as: task add TASK_HERE",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("You now have \"%s\" pending\n", task)
		db.CreateTask(task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
