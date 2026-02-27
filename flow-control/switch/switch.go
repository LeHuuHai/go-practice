package main

import "fmt"

func main() {
	// Không giống cpp, java..., break được tự động thêm vào cuối mỗi case,
	// do đó Go chỉ chạy case đùng mà bỏ qua tất cả các case sau đó
	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two") // Two
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// switch không có biểu thức -> kiểm tra điều kiện của case (<=> if else)
	switch {
	case x > 0:
		fmt.Println("Positive") // Positive
	case x < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
