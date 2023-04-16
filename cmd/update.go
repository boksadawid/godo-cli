package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/boksadawid/godo-cli/file_storage"
	"github.com/boksadawid/godo-cli/task_manager"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := file_storage.LoadFromJson()
		manager := task_manager.TaskManager{
			Tasks: tasks,
		}

		if err != nil {
			fmt.Println("Error initializing load:", err)
			os.Exit(1)
		}

		if len(args) < 6 {
			fmt.Println("Requires 6 arguments")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}

		deadline, err := time.Parse("2006-01-02 15:04:05", args[3])

		manager = *manager.Update(id, args[1], args[2], args[3], deadline, args[5] == "true", false)

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

var DoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark task as Done.",
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
		task := *manager.Get(id)
		manager = *manager.Update(id, task.Name, task.Description, task.Cron, task.Deadline, task.Backlog, true)

		if err != nil {
			fmt.Println("Error updating task:", err)
			os.Exit(1)
		}

		err = file_storage.SaveToJson(manager.Tasks)

		if err != nil {
			fmt.Println("Error saving task:", err)
			os.Exit(1)
		}
	},
}
