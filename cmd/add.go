/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// Task struct
type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

var taskFile = "tasks.json"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		tasks := readTasks()

		newTask := Task{
			ID:        len(tasks) + 1,
			Name:      taskName,
			CreatedAt: time.Now(),
		}
		tasks = append(tasks, newTask)

		saveTasks(tasks)
		fmt.Printf("Task added: %s\n", taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func readTasks() []Task {
	var tasks []Task
	file, err := os.ReadFile(taskFile)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
	return tasks
}

func saveTasks(tasks []Task) {
	file, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile(taskFile, file, 0644)
}
