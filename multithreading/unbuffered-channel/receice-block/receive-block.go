package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		x := 42
		fmt.Printf("Send %d\n", x)
		ch <- x
	}()

	fmt.Println("waiting...")
	v := <-ch
	fmt.Printf("Receive %d", v)
}
