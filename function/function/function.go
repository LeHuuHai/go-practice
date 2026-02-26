package main

import "fmt"

// Nếu 2 hay nhiều tham số có cùng kiểu,
// có thể bỏ qua dữ liệu ngoại trừ tham số cuối cùng
func add(a, b int) int {
	return a + b
}

// Hàm có thể gán vào biến, sử dụng làm đối số,
// hoặc được trả về từ hàm khác
func operator(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	fmt.Println(add(2, 3)) // 5

	// Biến kiểu hàm khi chưa khởi tạo là nil
	var f func(int, int) int
	f = add
	fmt.Println(operator(3, 3, f)) // 6

	// Hàm ẩn dnah định nghĩa tại thời điểm sử dụng,
	// có khả năng truy cập các biến bên trong hàm bao của nó,
	// trở thành closure
	a := 7
	func(val int) {
		fmt.Println(val + 1)
	}(a) // 8
}
