package main

import "fmt"

// Các giá trị trả về có thể được đặt tên,
// khi đó chúng được đăt như cú pháp tham số,
// và biểu thức return có thể không cần đối số.
func dosth(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	fmt.Println(dosth(5, 6)) // 11 -1
}
