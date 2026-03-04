package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // buffer size = 2
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("receive:", <-ch)
	}()

	ch <- 1 // không block
	fmt.Println("send 1 done")
	ch <- 2 // không block
	fmt.Println("send 2 done")
	ch <- 3 // block
	fmt.Println("send 3 done")
	time.Sleep(500 * time.Millisecond)
}
