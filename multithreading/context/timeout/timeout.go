package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Tạo một context với timeout 2 giây
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) { // goroutine này sẽ bị timeout sau 2 giây
		defer wg.Done()
		for i := 0; i < 100; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 1 timeout")
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
		fmt.Println("Task completed")
	}(ctx)

	wg.Wait()
	fmt.Println("Main goroutine done")
}
