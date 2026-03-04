package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("goroutine:", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println("main:", i)
		time.Sleep(200 * time.Millisecond)
	}
}
