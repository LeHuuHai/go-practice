package main

import (
	"fmt"
	"sync"
)

type struct1 struct {
	mu sync.Mutex
}

func main() {
	s1 := &struct1{}

	s1.mu.Lock()
	defer s1.mu.Unlock()
	// Some logic
	// ...
	// ...
	s1.mu.Lock() // Goroutine cố gắng lock lại mutex đã được lock -> deadlock
	defer s1.mu.Unlock()

	fmt.Println("Done")
}
