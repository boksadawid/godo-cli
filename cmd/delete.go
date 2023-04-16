package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/boksadawid/godo-cli/file_storage"
	"github.com/boksadawid/godo-cli/task_manager"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := file_storage.LoadFromJson()
		manager := task_manager.TaskManager{
			Tasks: tasks,
		}

		if err != nil {
			fmt.Println("Error initializing load:", err)
			os.Exit(1)
		}

		if len(args) == 0 {
			fmt.Println("Requires ID as argument")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])

		manager = *manager.Delete(id)

		if err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}

		err = file_storage.SaveToJson(manager.Tasks)

		if err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}
	},
}
