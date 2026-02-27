package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	v  int
}

// Tăng giá trị an toàn với nhiều goroutine
func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

// Lock và đọc giá trị
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func main() {

	var c Counter
	var wg sync.WaitGroup

	// Tạo 1000 goroutine tăng biến v
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	fmt.Println("final value:", c.Value())
}
