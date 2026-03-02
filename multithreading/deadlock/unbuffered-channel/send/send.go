package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1 // Main goroutine gửi dữ liệu vào channel nhưng không có goroutine nào nhận -> deadlock
	fmt.Println("Done")
}
