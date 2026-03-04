package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	mu      sync.RWMutex
}

// Đang Lock() mà RLock() thì deadlock vì RLock() chờ writer unlock
func ex1(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Ex1 begin")
	counter.mu.Lock()
	counter.mu.RLock()
	// do sth
	counter.mu.RUnlock()
	counter.mu.Unlock()
}

func ex2(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Ex2 begin")
	counter.mu.RLock()
	counter.mu.Lock()
	// do sth
	counter.mu.Unlock()
	counter.mu.RUnlock()
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	wg.Add(1)

	//go ex1(&counter, &wg)
	go ex2(&counter, &wg)

	wg.Wait()
	fmt.Println("main done")

}
