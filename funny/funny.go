package funny

import "fmt"

type nil int
type bool int8

// 预定义标识符，非关键字

// Types:
// bool byte complex64 complex128 error float32 float64
// int int8 int16 int32 int64 rune string
// uint uint8 uint16 uint32 uint64 uintptr
// Constants:
// true false iota
// Zero value:
// nil
// Functions:
// append cap close complex copy delete imag len
// make new panic print println real recover

func main()  {
	var i nil = 100
	var isSuc bool = 100
	fmt.Println(i)
	fmt.Println(isSuc)
}