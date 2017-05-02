package main


const (
	INT_MAX = 2147483647
)

/*
x * divisor = dividend
d + d + d = dividend

*/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend < 0 {

	}
	if divisor == 0 {
		panic("divide by zera")
	}
	result, step:=0, divisor
	for divisor<=dividend  {
		result++
		divisor+=step
	}
	return result
}
