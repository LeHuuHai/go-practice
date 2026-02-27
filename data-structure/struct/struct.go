package main

import "fmt"

type Vertex struct {
	X int
	Y int
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
}
