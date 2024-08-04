package main

import (
	"fmt"
	"time"
)

// Task represents a unit of work for the workers
type Task1 struct {
	id int
}

func main2() {
	//  Worker Pool: A set of worker goroutines that process tasks from a common channel.
	startTime := time.Now()
	totalJobs := 5

	for i := 1; i <= totalJobs; i++ {
		processJobs(Task1{id: i})
	}

	fmt.Println("Total time ", time.Since(startTime))
}

func processJobs(task Task1) {
	fmt.Println("Started  job", task.id)
	time.Sleep(time.Second)
	fmt.Println("Finished  job", task.id)
}
