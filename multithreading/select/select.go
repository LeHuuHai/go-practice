package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()
	go func() {
		select {
		case v := <-ch1:
			fmt.Printf("receive from channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("receive from channel 2: %d\n", v)
		}
	}()
	time.Sleep(5 * time.Second) // chờ goroutine chạy, nên dùng wait group
}
