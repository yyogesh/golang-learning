package main

import (
	"context"
	"fmt"
	"time"
)

func main4() {
	// Create a context with a 2-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// Simulate some work
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled due to timeout")
		fmt.Println(ctx.Err())
	}

}
