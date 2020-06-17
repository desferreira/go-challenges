package cmd

import (
	"fmt"
	"strconv"

	"github.com/desferreira/go-challenges/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark the todo as DONE",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Something went wrong with your args: ", err)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, _ := db.AllTasks()
		for _, v := range ids {
			if v <= 0 || v > len(tasks) {
				fmt.Println("Invalid task number")
				continue
			}
			task := tasks[v-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%v\" as complete\n", task.Key)
			} else {
				fmt.Printf("Task \"%v\" marked as complete", task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
