// main.go
package main

import (
	"os"
	"todo-cli/commands"
	"todo-cli/task"
)

func main() {
	var tasks []task.Task
	tasks, err := task.LoadTasks("tasks.json")
	if err != nil {
		tasks = []task.Task{}
	}
	commands.ExecuteCommand(os.Args[1:], &tasks)

}
