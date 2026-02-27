package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		ch <- id*100 + i
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	numWorkers := 3

	// Spawn workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	// Goroutine đặc biệt dùng để đóng channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Main đọc dữ liệu
	for v := range ch {
		fmt.Println("Main received:", v)
	}
	fmt.Println("All workers done")
}
