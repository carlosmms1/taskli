package task

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Description string `json:"description"`
	Done 		bool   `json:"done"`
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(fileName)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func Add(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, Task{Description: description})
	return saveTasks(tasks)
}

func List() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		status := " "
		if t.Done {
			status = "X"
		}
		fmt.Printf("[%s] %d - %s\n", status, i+1, t.Description)
	}
	return nil
}

func Done(indexStr string) error {
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return err
	}

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if (index < 1 || index > len(tasks)) {
		return fmt.Errorf("index overflow")
	}

	tasks[index - 1].Done = true
	return saveTasks(tasks)
}