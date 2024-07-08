package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

const taskFile = "task.json"

func main() {
	for {
		fmt.Println("Task List")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark Task as Complete")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')

		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			addTask()
		case "2":
			viewTasks()
		case "3":
			deleteTask()
		case "4":
			markTaskAsComplete()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again. Please enter a number between 1 and 5.")
		}
	}
}

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks file: %w", err)
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks from JSON: %w", err)
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal tasks to JSON: %w", err)
	}
	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks to file: %w", err)
	}
	return nil
}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	tasks = append(tasks, Task{Description: description, Completed: false})

	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}

	fmt.Println("Task added successfully.")
}

func viewTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Printf("%-5s %-30s %s\n", "ID", "Description", "Completed")
	fmt.Println("-----------------------------------------------")
	for i, task := range tasks {
		fmt.Printf("%-5d %-30s %t\n", i+1, task.Description, task.Completed)
	}
}

func deleteTask() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	viewTasks()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task number to delete: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskNum, err := strconv.Atoi(input)
	if err != nil || taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	// 5 [0, 1, 2, 3, 4]
	// tasks[:1] task[2:]
	tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)

	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}

	fmt.Println("Task deleted successfully.")
}

func markTaskAsComplete() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	viewTasks()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task number to mark as complete: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskNum, err := strconv.Atoi(input)
	if err != nil || taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks[taskNum-1].Completed = true

	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	fmt.Println("Task marked as complete.")
}
