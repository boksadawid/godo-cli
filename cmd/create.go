package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/boksadawid/godo-cli/file_storage"
	"github.com/boksadawid/godo-cli/task_manager"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new task.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := file_storage.LoadFromJson()
		manager := task_manager.TaskManager{
			Tasks: tasks,
		}

		if err != nil {
			fmt.Println("Error initializing load:", err)
			os.Exit(1)
		}

		if len(args) < 5 {
			fmt.Println("Requires 5 arguments")
			os.Exit(1)
		}

		deadline, err := time.Parse("2006-01-02 15:04:05", args[3])

		manager = *manager.Add(args[0], args[1], args[2], deadline, args[4] == "true")

		if err != nil {
			fmt.Println("Error creating task:", err)
			os.Exit(1)
		}

		err = file_storage.SaveToJson(manager.Tasks)

		if err != nil {
			fmt.Println("Error saving task:", err)
			os.Exit(1)
		}
	},
}
