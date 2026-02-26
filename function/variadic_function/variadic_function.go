package main

import "fmt"

// Hàm đa biến cho phép số lượng tham số tùy ý (0..n).
// Tập các tham số này được coi như 1 slice
// Tập các tham số phải đứng cuối cùng trong ds tham số
func sum(nums ...int) int {
	ans := 0
	for _, n := range nums {
		ans += n
	}
	return ans
}

func main() {
	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
	fmt.Println(sum())              // 0
}
