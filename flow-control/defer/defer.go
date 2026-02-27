package main

import "fmt"

// Defer hoãn việc gọi hàm đến khi hàm hiện tại kết thúc.
// Nếu nhiều defer -> chạy theo thứ tự LIFO.
// Tham số được tính theo thời điểm defer thay vì thời điểm thực thi hàm
func ex1() {
	a := 10
	defer fmt.Println(a)
	a = 20
}

func ex2() {
	a := 10
	defer func(v int) {
		fmt.Println(v)
	}(a)
	a = 20
}

func ex3() {
	a := 10
	defer func() {
		fmt.Println(a)
	}()
	a = 20
}

func main() {
	// hàm ex1() và ex2() tại dòng defer: tạo bản sao đối số a và con trỏ hàm defer
	// -> khi thực thi dùng bản copy của a tại thời điểm defer
	ex1() // 10
	ex2() // 10
	// hàm ex3() tại dòng defer: chỉ có con trỏ hàm được lưu vào stack, còn closure tham chiếu đối số a
	// -> khi thực thi dùng đối số a có giá trị đã thay đổi
	ex3() // 20
}
