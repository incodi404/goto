package main

import (
	"fmt"

	"github.com/incodi404/goto/internal/workflow"
)

func main() {
	// fmt.Println("Welcome to Goto! A Golang based Workflow Automation Engine")
	wf := workflow.Workflow{
		ID:     "wf-1",
		Name:   "Test Workflow",
		Status: "Pending",
		Tasks: []workflow.Task{
			{ID: "t1", Type: "email"},
			{ID: "t1", Type: "email"},
		},
	}

	// fmt.Println("Workflow: ", wf.Tasks)
	for index, task := range wf.Tasks {
		fmt.Println("Executing task: ", task.ID, task.Type, index)
	}

	for i := 0; i <= 5; i++ {
		fmt.Println("Index: ", i)
	}
}
