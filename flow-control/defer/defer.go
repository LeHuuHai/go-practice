package main

import "fmt"

func ex1() {
	a := 10
	defer fmt.Println(a)
	a = 20
}

func ex2() {
	a := 10
	defer func(v int) {
		fmt.Println(v)
	}(a)
	a = 20
}

func ex3() {
	a := 10
	defer func() {
		fmt.Println(a)
	}()
	a = 20
}

func main() {
	ex1() // 10
	ex2() // 10
	ex3() // 20
}
