package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/boksadawid/godo-cli/file_storage"
	"github.com/boksadawid/godo-cli/task_manager"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/spf13/cobra"
)

var listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "List all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := file_storage.LoadFromJson()

		if err != nil {
			fmt.Println("Error listing tasks:", err)
			os.Exit(1)
		}

		printTasksTable(tasks)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List not done tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := file_storage.LoadFromJson()
		manager := task_manager.TaskManager{
			Tasks: tasks,
		}

		tasks = manager.ListNotDone()

		if err != nil {
			fmt.Println("Error listing tasks:", err)
			os.Exit(1)
		}

		printTasksTable(tasks)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Requires ID as argument")
			os.Exit(1)
		}

		tasks, err := file_storage.LoadFromJson()

		if err != nil {
			fmt.Println("Error loading storage:", err)
			os.Exit(1)
		}

		manager := task_manager.TaskManager{
			Tasks: tasks,
		}

		id, err := strconv.Atoi(args[0])

		task := manager.Get(id)
		task_list := []task_manager.Task{*task}

		if err != nil {
			fmt.Println("Error getting task:", err)
			os.Exit(1)
		}

		printTasksTable(task_list)
	},
}

func printTasksTable(tasks []task_manager.Task) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Name", "Description", "Cron", "Deadline", "Backlog", "Done")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, task := range tasks {
		tbl.AddRow(
			task.ID, task.Name, task.Description, task.Cron, task.Deadline.Format("2006-01-02 15:04:05"), task.Backlog, task.Done)
	}

	tbl.Print()
}
