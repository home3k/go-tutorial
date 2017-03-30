package main

import (
	"fmt"
	"strconv"
)

type Weekday int

const (
	Sun Weekday = iota
	Mon
	Tues
	Wed
	Thur
	Fri
	sAT
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	w := 1 + 2i
	v := 3 + 4i
	fmt.Println(w * v)

	ai := 19
	fmt.Println(strconv.Itoa(ai))

	ia := "23"
	fmt.Println(strconv.Atoi(ia))

	var ar [3]int
	fmt.Println(ar[0])

	var arr [3]int = [3]int{1, 3, 4}
	for index, value := range arr {
		fmt.Printf("index:%d, value: %d\n", index, value)
	}
	fmt.Println("------")
	for _, value := range arr {
		fmt.Printf("value: %d\n", value)
	}

	q := [...]int{10,20,30}
	fmt.Printf("type %T", q)
}
