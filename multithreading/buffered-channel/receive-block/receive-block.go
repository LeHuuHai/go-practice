package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("send:", 42)
		ch <- 42
	}()

	fmt.Println("waiting...")
	v := <-ch // block vì buffer rỗng
	fmt.Println("received:", v)
	time.Sleep(500 * time.Millisecond)
}
