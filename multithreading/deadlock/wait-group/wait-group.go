package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() // Main goroutine chờ đợi nhưng không có goroutine nào gọi wg.Done() -> deadlock
	// some logic
	// ...
	fmt.Println("Done")
}
