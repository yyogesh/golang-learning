package main

import (
	"context"
	"fmt"
	"time"
)

func main2() {
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

func work(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Work canceled")
		fmt.Println(ctx.Err())
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed")

	}
}
