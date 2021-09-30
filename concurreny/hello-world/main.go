package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go goroutine(ctx)

	time.Sleep(5*time.Second)
	cancel()

	fmt.Println("bye")
	time.Sleep(5*time.Second)
}

func goroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("hello")
		}
	}
}

