package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define a "Task" type
type Task struct {
	ID int
	Title string
	Done bool
}

// Temporary main function to test
func main() {
	tasks := []Task {
		{ID: 1, Title: "Test 1", Done: false},
		{ID: 2, Title: "Test 2", Done: false},
	}

	// Add Task
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a new task: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	newTask := Task {
		ID: len(tasks) + 1,
		Title: input,
		Done: false,
	}

	// Append
	tasks = append(tasks, newTask)

	// Ask which one is completed
	fmt.Print("Enter the ID of the task to mark as done: ")
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput)
	id, err := strconv.Atoi(idInput) // Convert string to int

	if err != nil {
		fmt.Println("Invalid ID input.")
		return
	}

	// Mark task done
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			break
		}
	}

	// Print updated task
	fmt.Println("\nYour updated task list:")
	for _, task := range tasks {
		status := " "

		if task.Done {
			status = "Y"
		}

		fmt.Println("- [%v] %s\n", task.Done, status, task.Title)
	}

	// Save task to file
	file, err := os.Create("tasks.json")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)

	if err != nil {
		fmt.Println("Error writing JSON:", err)
	}
}