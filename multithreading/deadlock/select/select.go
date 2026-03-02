package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Cả 2 case đều không ready -> deadlock
	select {
	case <-ch1:
		fmt.Println("case 1 ready")
	case ch2 <- 1:
		fmt.Println("case 2 ready")
	}

	fmt.Println("Done")
}
