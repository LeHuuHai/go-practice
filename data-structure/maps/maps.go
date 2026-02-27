package main

import "fmt"

func change(mp map[string]int) {
	mp["a"] = 999
}

func main() {

	// Map là struct chứa con trỏ đến underlying data là hmap
	// Key phải comparable

	// zero value của map là nil
	var m map[string]int
	fmt.Println(m == nil) // true

	// Khởi tạo map bằng make
	m = make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	fmt.Println(m) // map[a:10 b:20]

	// Thay đổi element
	change(m)
	fmt.Println(m) // map[a:999 b:20]

	// Kiểm tra key có tồn tại, nếu không, trả về zero value và false
	v, ok := m["x"]
	fmt.Println(v, ok) // 0 false

	// Xóa phần tử
	delete(m, "b")
	fmt.Println(m) // map[a:999]

	// Thứ tự duyệt map là ngẫu nhiên
	m["c"] = 100
	for k, v := range m {
		fmt.Printf("(%s, %d) ", k, v) // (a, 999) (c, 100)
	}
}
