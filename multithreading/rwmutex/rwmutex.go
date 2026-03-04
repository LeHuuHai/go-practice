package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.RWMutex
}

func Read(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.mu.RLock()
	fmt.Printf("read %d\n", counter.value)
	counter.mu.RUnlock()
}

func Write(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.mu.Lock()
	counter.value++
	fmt.Println("added")
	counter.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(20)
	counter := Counter{}

	go func() {
		for i := 0; i < 10; i++ {
			go Read(&counter, &wg)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			go Write(&counter, &wg)
		}
	}()

	wg.Wait()
}
