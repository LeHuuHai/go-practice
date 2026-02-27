package main

import "fmt"

func main() {
	// if có thể bao gồm 1 biểu thức ngắn để thực thi trước
	// biểu thức điều kiện, các biến được khai báo trong biếu
	// thức này có thể được truy cập trong khối if, else của nó
	if num := 8; num%2 == 0 {
		fmt.Println("Even:", num) // Even: 8
	} else {
		fmt.Println("Odd:", num)
	}
}
