// commands/commands.go
package commands

import (
	"flag"
	"fmt"
	"os"
	"todo-cli/task" // Adjust to your module path
)

// ExecuteCommand parses the command and executes the appropriate action
func ExecuteCommand(args []string, tasks *[]task.Task) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Define flags
	//taskName := addCmd.String("name", "", "Name of the task")
	filterDone := listCmd.Bool("done", false, "Filter tasks by done status")
	sortBy := listCmd.String("sort", "", "Sort tasks by date added")

	if len(args) < 1 {
		fmt.Println("expected 'add', 'done', 'list', or 'delete' subcommands")
		os.Exit(1)
	}

	switch args[0] {
	case "list":
		listCmd.Parse(args[1:])
		task.ListTasks(*tasks, *filterDone, *sortBy)
	case "add":
		addCmd.Parse(args[1:])
		if addCmd.NArg() < 1 {
			fmt.Println("Please provide a task to add")
			os.Exit(1)
		}
		taskName := addCmd.Arg(0)
		task.AddTask(taskName, tasks)
	case "done":
		doneCmd.Parse(args[1:])
		if doneCmd.NArg() < 1 {
			fmt.Println("Please provide a task ID to mark as done")
			os.Exit(1)
		}
		taskID := doneCmd.Arg(0)
		task.MarkDone(taskID, tasks)
	case "delete":
		deleteCmd.Parse(args[1:])
		if deleteCmd.NArg() < 1 {
			fmt.Println("Please provide a task ID to delete")
			os.Exit(1)
		}
		taskID := deleteCmd.Arg(0)
		task.DeleteTask(taskID, tasks)
	default:
		fmt.Println("expected 'add', 'done', or 'delete' subcommands")
		os.Exit(1)
	}
}
