package main

import (
	"errors"
	"fmt"
)

// go dùng error như 1 giá trị thường để kiểm soát lỗi
// khi 1 hàm có thể xảy ra lỗi, cần bao gồm error trong
// danh sách trả về, error theo qy ước nằm ở cuối danh
// sách trả về. Nếu trả về lỗi, các giá trị không an toàn
// (là các giá trị không có ý nghĩa khi lỗi) nên được trả về zero
func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide for 0")
	}
	return float64(a) / float64(b), nil
}

// custom error type
type MyError struct {
	Code int
	Msg  string
}

// impl bằng cách thêm triển khai hàm Error()
func (e *MyError) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Msg)
}

func main() {
	// ở hàm gọi, luôn kiểm tra có lỗi hay không trước khi sử dụng kết quả trả về
	res, err := divide(5, 0)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("cannot divide")
	}

	// thử custom err
	res, err = func() (float64, error) {
		return 0, &MyError{100, "some error msg"}
	}()
	if e, ok := err.(*MyError); ok {
		fmt.Printf("code=%d, msg=%s\n", e.Code, e.Msg) // code=100, msg=some error msg
	}

}
