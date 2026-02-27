package main

import (
	"fmt"
	"time"
)

// channel để giao tiếp giữa các goroutine

// unbuffered channel không có buffer, dữ liệu được truyền khi sender và receiver gặp nhau.
// Thao tác gửi và nhận đảm bảo cùng hoàn thành
func ex1() {
	// tạo 1 unbuffered channel có thể truyền dữ liệu kiểu int
	ch1 := make(chan int)
	go func() {
		ch1 <- 10 // block cho đến khi có receiver
	}()
	go func() {
		v := <-ch1 // block cho đến khi có sender
		fmt.Printf("receive %d", v)
	}()
	time.Sleep(5 * time.Second) // cho goroutine chạy, có thể dùng wait group
}

// buffered channel có mảng vòng làm buffer là trung gian. Thao tác gửi sẽ block khi buffer
// đầy và không có receiver đang chờ. Thao tác nhận sẽ block khi buffer rỗng và không có sender đang chờ.
// Nếu có goroutine chờ ở phía đối diện (trong sendq hay receiveq), runtime có thể truyền dữ liệu trực tiếp
// mà không sử dụng buffer.
func ex2() {
	// tạo 1 buffer channel có kích thước là 5
	ch1 := make(chan int, 5)
	go func() {
		ch1 <- 10
		ch1 <- 20
	}()
	go func() {
		v := <-ch1
		fmt.Printf("receive %d\n", v)
		v = <-ch1
		fmt.Printf("receive %d\n", v)
	}()
	time.Sleep(5 * time.Second) // chờ goroutine chạy, có thể dùng wait group
}

// có thể đóng channel bằng hàm close, và chỉ nên đóng ở sender, khi này thao tác gửi sẽ gây panic.
// nếu còn dữ liệu và channel đã close -> trả về dữ liệu và true; nếu hết dữ liệu và đã đóng -> trả về zero và false
func ex3() {
	ch := make(chan int, 1)
	ch <- 10
	close(ch)

	v, ok := <-ch
	fmt.Println(v, ok) // 10 true

	v, ok = <-ch
	fmt.Println(v, ok) // 0 false
}

// dùng vòng lặp for range để đọc dữ liệu từ channel
func ex4() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch { // tương đương v, ok := <-ch; if !ok break
		fmt.Printf("receice %d\n", v)
	}
}

func main() {
	ex4()
}
