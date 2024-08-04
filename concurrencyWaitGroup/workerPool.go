package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work for the workers
type Task struct {
	id int
}

func main() {
	startTime := time.Now()
	const numWorkers = 2
	const numTasks = 5

	tasks := make(chan Task, numTasks)
	var wg sync.WaitGroup

	// Launch worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Send tasks to the task channel
	for j := 1; j <= numTasks; j++ {
		tasks <- Task{id: j}
	}

	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("Total time ", time.Since(startTime))
}

// worker is the function that each worker goroutine will run
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task.id)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished task %d\n", id, task.id)
	}
}
