package main

import (
	"fmt"
	"os"

	"github.com/carlosmms1/taskli/internal/task"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Use todo add|list|done <n>")
		return
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("You need to input task description...")
			return
		}
		err := task.Add(args[2])
		if err != nil {
			fmt.Println("There was an error: ", err)
			return
		}
		fmt.Println("Task created successfully!")

	case "list":
		err := task.List()
		if err != nil {
			fmt.Println("There was an error listing the tasks:", err)
		}
	
	case "done":
		if len(args) < 3 {
			fmt.Println("You need to insert the task index...")
			return
		}
		err := task.Done(args[2])
		if err != nil {
			fmt.Println("There was an error to mark done:", err)
			return
		}
		fmt.Println("Task marked as done successfully!")

	default:
		fmt.Println("Unknown option...")
	}
}
