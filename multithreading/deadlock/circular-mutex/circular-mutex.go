package main

import (
	"fmt"
	"sync"
	"time"
)

type struct1 struct {
	mu sync.Mutex
}

type struct2 struct {
	mu sync.Mutex
}

func main() {
	s1 := &struct1{}
	s2 := &struct2{}
	var wg sync.WaitGroup
	wg.Add(2)

	// goroutine 1 chờ lock từ goroutine 2, goroutine 2 chờ lock từ goroutine 1 -> deadlock
	go func(a *struct1, b *struct2) {
		a.mu.Lock()
		defer a.mu.Unlock()
		time.Sleep(2 * time.Second)
		b.mu.Lock()
		defer b.mu.Unlock()
		wg.Done()
	}(s1, s2)

	go func(a *struct1, b *struct2) {
		b.mu.Lock()
		defer b.mu.Unlock()
		time.Sleep(2 * time.Second)
		a.mu.Lock()
		defer a.mu.Unlock()
		wg.Done()
	}(s1, s2)

	wg.Wait()
	fmt.Println("Done")
}
