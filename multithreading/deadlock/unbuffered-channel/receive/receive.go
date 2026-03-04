package main

import "fmt"

func main() {
	ch := make(chan int)
	<-ch // Main goroutine chờ nhận từ channel nhưng không có goroutine nào gửi -> deadlock
	fmt.Println("Done")
}
