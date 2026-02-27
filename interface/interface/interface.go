package main

import "fmt"

// trong go không có từ khóa inplements. Bất kể kiểu dữ liệu nào định nghĩa
// các phương thức của interface đều tự động implement interface đó

type Speaker interface {
	Speak()
}

type Cat struct {
}

// impl
func (cat *Cat) Speak() {
	fmt.Println("meow meow")
}

func Play(s Speaker) {
	s.Speak()
}

func main() {
	cat := Cat{}
	Play(&cat) // meow meow, *Cat đã impl Speaker
}
