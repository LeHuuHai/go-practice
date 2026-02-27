package main

import "fmt"

// array là value type, chỉ thay đổi bản copy
func change(a [3]int) {
	a[0] = 999
}

// truyền con trỏ mới thay đổi mảng gốc
func changePtr(a *[3]int) {
	a[0] = 999
}

func main() {

	// array phải có kích thước cố định
	var a [3]int = [3]int{1, 2, 3}

	change(a)
	fmt.Println("after change:", a) // after change: [1 2 3]

	changePtr(&a)
	fmt.Println("after changePtr:", a) // after changePtr: [999 2 3]

	// So sánh trực tiếp
	c := [3]int{999, 2, 3}
	fmt.Println("a == c:", a == c) // a == c: true
}
