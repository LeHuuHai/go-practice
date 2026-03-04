package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Tạo một context với timeout 2 giây
	childctx, _ := context.WithCancel(ctx)                                  // Tạo một context con từ context cha

	go func(ctx context.Context) {
		<-ctx.Done() // Goroutine này sẽ chờ cho đến khi context con bị hủy hoặc timeout
		fmt.Println("Child goroutine nhận được tín hiệu hủy")
	}(childctx)

	cancel() // Hủy context cha, điều này sẽ tự động hủy context con

	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine done")
}
