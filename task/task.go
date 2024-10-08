package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mergestat/timediff"
)

// Task struct defines the properties of a task
type Task struct {
	ID        int
	Name      string
	Done      bool
	Timestamp int64
}

// AddTask adds a new task to the list
func AddTask(name string, tasks *[]Task) {
	lastID := 0
	for _, task := range *tasks {
		lastID = max(lastID, task.ID)
		if task.ID > lastID {
			lastID = task.ID
		}
	}
	task := Task{
		ID:        lastID + 1,
		Name:      name,
		Done:      false,
		Timestamp: time.Now().Unix(),
	}
	*tasks = append(*tasks, task)
	SaveTasks(*tasks, "tasks.json")
}

// ListTasks lists all tasks
func ListTasks(tasks []Task, filterDone bool, sortBy string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	// for _, task := range tasks {
	// 	status := " "
	// 	if task.Done {
	// 		status = "x"
	// 	}
	// 	fmt.Printf("[%s] %s %d: %s\n", status, timediff.TimeDiff(time.Unix(task.Timestamp, 0)), task.ID, task.Name)
	// }
	tasksCopy := make([]Task, len(tasks))
	copy(tasksCopy, tasks)
	if sortBy == "earliest" {
		sort.Slice(tasksCopy, func(i, j int) bool {
			return tasksCopy[i].Timestamp < tasksCopy[j].Timestamp
		})
	} else if sortBy == "latest" {
		sort.Slice(tasksCopy, func(i, j int) bool {
			return tasksCopy[i].Timestamp > tasksCopy[j].Timestamp
		})
	} else if sortBy != "" {
		fmt.Println("Invalid sortBy flag. Please use 'earliest' or 'latest'")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"#", "", "Task Name", "Date Added"})
	for _, task := range tasksCopy {
		status := " "
		if filterDone && !task.Done {
			continue
		}
		if task.Done {
			status = "x"
		}
		t.AppendRow([]interface{}{task.ID, status, task.Name, timediff.TimeDiff(time.Unix(task.Timestamp, 0))})
	}
	t.Render()
}

// MarkDone marks a task as done by its ID
func MarkDone(idStr string, tasks *[]Task) {
	id, _ := strconv.Atoi(idStr)
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Done = true
			break
		}
	}
	SaveTasks(*tasks, "tasks.json")
}

// DeleteTask deletes a task by its ID
func DeleteTask(idStr string, tasks *[]Task) {
	id, _ := strconv.Atoi(idStr)
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			break
		}
	}
	SaveTasks(*tasks, "tasks.json")
}

// SaveTasks saves the tasks to a file in JSON format
func SaveTasks(tasks []Task, filename string) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// LoadTasks loads tasks from a JSON file
func LoadTasks(filename string) ([]Task, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
