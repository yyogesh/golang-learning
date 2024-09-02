package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a 2-second timeout
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	// Simulate some work
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled due to deadline")
		fmt.Println(ctx.Err())
	}

}
