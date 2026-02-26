package main

import "fmt"

// Trả về hàm sử dụng biến local cnt
// Biến này không bị thu hồi mà được escape vào heap
// và tồn tại cùng với hàm trả về dù hàm dosth đã kết thúc
func dosth() func() int {
	cnt := 0
	return func() int {
		cnt += 1
		return cnt
	}
}

func main() {
	f := dosth()
	// Biến cnt vẫn tồn tại, và được tăng lên mỗi lần gọi f()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2

	// Biến local được sử dụng bởi closure không phải là bản
	// copy dùng riêng cho closure mà là tham chiếu thực sự
	a := 0
	f2 := func() int {
		a += 1
		return a
	}
	fmt.Println(f2()) // 1
	fmt.Println(a)    // 1

	// Vì là tham chiếu thực sự nên nếu capture biến đếm trong vòng lặp
	// thì có thể dẫn đến lỗi
	funcs := make([]func(), 0)
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for _, f := range funcs {
		f() // 3, 3, 3 (ver < 1.22)
	}
}
