package main

import "fmt"

const (
	INT_MAX = 2147483647
	INT_MIN = -2147483648
)

/*

x * divisor = dividend
3 10
4 * 3 > 10
3 * 3 < 10

d + d + d = dividend

*/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 0 {
		panic("divide by zera")
	}
	sign, step := 1, 1
	if dividend < 0 {
		dividend = -dividend
		sign = -sign
	}
	if divisor < 0 {
		divisor = -divisor
		sign = -sign
	}
	if dividend == divisor {
		return sign
	}
	td := divisor
	for td < dividend {
		step<<=1
		td <<= 1
	}
	if step == 1 {
		return 0
	}
	for td > dividend {
		td -= divisor
		step--
	}
	result := sign * step
	if result > INT_MAX {
		return INT_MAX
	} else if result < INT_MIN {
		return INT_MIN
	} else {
		return result
	}
}

func main() {
	fmt.Println(divide(-2147483648, 1))
	fmt.Println(divide(-2147483648 / 2 + 1, 1))
	fmt.Println(divide(1, 1))
	fmt.Println(divide(2, 1))
	fmt.Println(divide(1, 2))
	fmt.Println(divide(10, 3))
}
