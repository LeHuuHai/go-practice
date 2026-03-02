package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		<-ch1
		ch2 <- 1
		wg.Done()
	}()

	go func() {
		<-ch2
		ch1 <- 1
		wg.Done()
	}()

	// Wait cho 2 goroutine hoàn thành nhưng cả 2 đều đang chờ nhau -> deadlock
	wg.Wait()
	fmt.Println("Done")
}
