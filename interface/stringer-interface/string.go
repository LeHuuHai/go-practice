package main

import (
	"fmt"
)

// interface stringer chứa phương thức String() cho phép định nghĩa định dạng in của cấu trúc dữ liệu
// type Stringer interface {
// 	String() string
// }

type Emp struct {
	Name string
	Age  int
}

func (e Emp) String() string {
	return fmt.Sprintf("His name is %s. He is %d years old.", e.Name, e.Age)
}

func main() {
	emp := Emp{"Hari", 23}
	fmt.Println(emp) // "His name is Hari. He is 23 years old." thay vì "{Hari 23}"

}
