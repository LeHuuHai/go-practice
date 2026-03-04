package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// embedd struct
type Circle struct {
	Vertex // embedded struct
	R      float64
}

// embedded struct có trường trùng tên
type Other struct {
	X int
	A Vertex
}

type Edge struct {
	A, B Vertex
}

func main() {
	v1 := Vertex{1, 2}            // has type Vertex
	v2 := Vertex{X: 1}            // Y:0 is implicit
	v3 := Vertex{}                // X:0 and Y:0
	p := &Vertex{1, 2}            // has type *Vertex
	q := new(Vertex)              // locate memory, zero value, <=> q:=&Vertex{}
	fmt.Println(v1, v2, v3, p, q) // {1 2} {1 0} {0 0} &{1 2} &{0 0}

	// trong go cho phép truy cập field struct mà không cần dereference
	fmt.Println(p.X) // 1

	// các trường của struct embedded được truy cập trực tiếp thông qua tên trường,
	// không cần thông qua tên struct
	c := Circle{Vertex{0, 0}, 5} // has type Circle
	fmt.Println(c.X, c.Y, c.R)

	// nếu trường trùng tên, ưu tiên truy cập trường của struct hiện tại hoặc của struct embedded gần nhất,
	// nếu muốn truy cập trường của struct embedded thì phải truy cập thông qua tên struct để phân biệt
	d := Other{X: 1, A: Vertex{2, 3}}
	fmt.Println(d.X, d.A.X)

	// nếu nhiều trường trùng tên ở cùng một mức embedd,
	// thì phải truy cập thông qua tên struct để phân biệt
	e := Edge{A: Vertex{1, 2}, B: Vertex{3, 4}}
	fmt.Println(e.A.X, e.B.X)
}
