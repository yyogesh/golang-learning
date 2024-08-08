package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	ch := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := job()
				ch <- result
			}()
		}
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Reading channel value ", v)
	}
	fmt.Println("Total time ", time.Since(startTime))
}

func job() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
