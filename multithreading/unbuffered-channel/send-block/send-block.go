package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("waiting...")
		ch <- 42
		fmt.Println("goroutine: sent")
	}()

	time.Sleep(2 * time.Second)
	v := <-ch
	fmt.Println("main received:", v)
	time.Sleep(500 * time.Millisecond)
}
