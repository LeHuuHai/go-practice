package main

import "fmt"

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Msg)
}

func errorable() (int, error) {
	var err *MyError = nil
	return 0, err
}

func main() {
	res, err := errorable()
	if err == nil {
		fmt.Println("no error, res =", res)
	} else { // err là interface error {type: MyError, value=nil} != nil => sai logic
		// Best practice nên trả về nil trực tiếp thay vì wrap vào interface
		fmt.Println("error:", err)
	}
}
