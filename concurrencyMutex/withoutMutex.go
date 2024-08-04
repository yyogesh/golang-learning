package main

import (
	"fmt"
	"sync"
)

var x1 = 0

func main1() {
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment1(&w)
	}

	w.Wait()

	fmt.Println("final value of x", x) // 1000
}

func increment1(wg *sync.WaitGroup) {
	x1++
	wg.Done()
}
