package main

import "fmt"

// Hàm cho phép trả về nhiều giá trị
// Chú ý chữ ký hàm
func dosth(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	fmt.Println(dosth(5, 6)) // 11 -1
}
