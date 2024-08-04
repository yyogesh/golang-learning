package main

import "fmt"

func printFunc(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

// close channel
func main7() {
	ch := make(chan int)
	go printFunc(ch)

	// for {
	// 	v, ok := <-ch
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Println("Received ", v, ok)
	// }

	for v := range ch {
		fmt.Println("Received ", v)
	}
}
