package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	// WaitGroup: Keeps track of a set of goroutines and waits for them to finish.
	// Worker Pool: A set of worker goroutines that process tasks from a common channel.

	no := 5

	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All go routines finished")

}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // Signal to the WaitGroup that this goroutine has completed its task.
}
