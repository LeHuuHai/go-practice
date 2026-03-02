package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6 // Main goroutine gửi dữ liệu vào channel nhưng channel đã đầy -> deadlock
	fmt.Println("Done")
}
