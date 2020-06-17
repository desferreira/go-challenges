package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/desferreira/go-challenges/task/cmd"
	"github.com/desferreira/go-challenges/task/db"
)

func main() {
	dbPath := filepath.Join("./", "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

/*
Handle errors in application
*/
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
