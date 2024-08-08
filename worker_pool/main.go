package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a festival task
type Task struct {
	ID   int
	Name string
}

// Result represents the outcome of a task
type Result struct {
	TaskID int
	Output string
}

func main() {
	fmt.Println("Organizing the School Festival...")

	tasks := []Task{
		{1, "Decorate the hall"},
		{2, "Prepare welcome drinks"},
		{3, "Set up the stage"},
		{4, "Arrange chairs"},
		{5, "Test sound system"},
		{6, "Hang banners"},
		{7, "Prepare food stalls"},
		{8, "Set up game booths"},
		{9, "Coordinate with performers"},
		{10, "Prepare goodie bags"},
	}

	numWorkers := 3

	tasksChan := make(chan Task, len(tasks))
	resultsChan := make(chan Result, len(tasks))

	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasksChan, resultsChan, &wg)
	}

	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)

	// Start a goroutine to close results channel after all workers finish
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Print the results

	for result := range resultsChan {
		fmt.Printf("Task %d completed: %s\n", result.TaskID, result.Output)
	}

	fmt.Println("All tasks are completed! The School Festival is ready to begin!")

}

func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	// chan   // read-write
	// <-chan // read-only
	// chan<- // write-only
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d started task: %s\n", id, task.Name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate work
		output := fmt.Sprintf("Worker %d finished '%s'", id, task.Name)
		results <- Result{task.ID, output}
	}

}
