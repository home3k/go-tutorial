package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	fmt.Println(unsafe.Sizeof(float64(0)))

	var x struct{
		a bool
		b int16
		c []int
	}
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x))+unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)
}
