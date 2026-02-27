package main

import "fmt"

type Emp struct {
	Name   string
	Age    int
	Salary int
}

// Hàm có tham số receiver là value
func (e Emp) rename(newName string) {
	e.Name = newName
}

// Hàm có tham số receiver là pointer
func (e *Emp) changeSalary(newSalary int) {
	e.Salary = newSalary
}

func main() {
	e := Emp{"Hari", 22, 1000}
	// Chỉ thành viên của receiver mới gọi được phương thức
	// tham số receiver là giá trị -> tạo bản sao -> không thay đổi đối số truyền vào
	e.rename("John")
	// tham số receiver là pointer -> có thể thay đổi giá trị của struct được trỏ đến
	e.changeSalary(2000)
	fmt.Println(e) // {Hari 22 2000}

	// chỉ quan tâm kiểu tham số receiver là value hay pointer, đối số của nó sẽ tự được convert tương ứng
	// truyền pointer vào hàm nhận value -> tự dereference
	ptr := &Emp{"Abel", 23, 1500}
	ptr.rename("John")
	fmt.Println(ptr) // &{Abel 23 2000}

	// truyền value vào hàm nhận pointer -> tự lấy địa chỉ &
	val := Emp{"Bob", 23, 3000}
	val.changeSalary(4000)
	fmt.Println(val) // {Bob 23 4000}
}
