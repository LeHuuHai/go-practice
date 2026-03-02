package main

import "fmt"

func produce(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go produce(ch)
	}

	// Main đọc dữ liệu trong for range nhưng không có goroutine nào đóng channel -> deadlock
	for v := range ch {
		fmt.Printf("receive %d\n", v)
	}

	fmt.Println("Done")
}
