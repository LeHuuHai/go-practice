package main

import "fmt"

func main() {
	// Không yêu cầu các biểu thức init, condition, post nằm trong (),
	// nhưng yêu cầu thân vòng for phải nằm trong {} và { nằm cùng dòng for
	//  do cơ chế tự đánh dấu ;
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10

	// các biểu thức init và post là optional, khi này for giống như
	// while trong các ngôn ngữ cpp, java
	cnt := 0
	mul := 0
	for cnt < 5 { // for condition {}
		mul += 2
		cnt++
	}
	fmt.Println(mul) // 10

}
