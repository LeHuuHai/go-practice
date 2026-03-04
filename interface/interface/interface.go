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

// embedd interface
type Mover interface {
	Move()
}

// interface này chứa chữ kí của cả 2 phương thức
type Animal interface {
	Speaker
	Mover
}

type Dog struct {
}

// impl
func (d *Dog) Speak() {
	fmt.Println("woof woof")
}

func (d *Dog) Move() {
	fmt.Println("Dog is running")
}

func HaveAPet(a Animal) {
	a.Speak()
	a.Move()
}

func main() {
	cat := Cat{}
	Play(&cat) // meow meow, *Cat đã impl Speaker

	dog := Dog{}
	HaveAPet(&dog) // *Dog đã impl Animal
}
