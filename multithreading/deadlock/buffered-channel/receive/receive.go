package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	<-ch // Main goroutine chờ nhận từ channel nhưng channel rỗng -> deadlock
	fmt.Println("Done")
}
