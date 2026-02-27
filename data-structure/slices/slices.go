package main

import "fmt"

func main() {
	// Slice là view động trên 1 mảng. Nó là struct chứa con trỏ đến mảng mà
	// nó quan sát, độ dài và sức chứa của nó
	// Nó là reference type nên có thể nil.

	// slice hình thành bằng cách chỉ định hai số giới hạn low hight của array, giá trị mặc định là 0 và arr_len
	// length: số lượng element của slice (high-low)
	// capacity: số lượng element của array tính từ element đầu tiên của slice (arr_len - low)
	arr := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[1:5]  // [1,5)
	fmt.Println(s) // [1 2 3 4]

	// Sự thay đổi của element trong slice dẫn đến sự thay đổi trong array và ngược lại.
	// Các slice khác cũng có thể thấy sự thay đổi này
	p := arr[2:4]    // [2,4)
	fmt.Println(p)   // [2 3]
	p[0] = 0         // p[0] <=> arr[2]
	fmt.Println(p)   // [0 3]
	fmt.Println(s)   // [1 0 3 4]
	fmt.Println(arr) // [0 1 0 3 4 5 6 7]

	// Hàm make cấp phát mảng zero value và trả về 1 slice tham chiếu đến mảng đó
	a := make([]int, 5)
	b := make([]int, 0, 5)
	fmt.Println(a) // [0 0 0 0 0]
	fmt.Println(b) // []

	// Hàm append thêm phần tử mới vào slice. Nếu đủ capacity, thay đổi giá trị trong array và
	// slice quan sát thêm phần tử này, nếu không, tạo 1 array mới với mở rộng từ array cũ
	// và trả về slice mới quan sát array mới này
	arr2 := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	c := arr2[5:8] // [5 6 7]
	c = append(c, 0)
	fmt.Println(c) // [5 6 7 0]
	d := append(c, 0)
	fmt.Println(d) // [5 6 7 0 0], rellocate
}
