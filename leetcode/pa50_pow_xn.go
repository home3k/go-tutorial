package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := myPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else if n > 0 {
		return x * half * half
	} else {
		return 1 / x * half * half
	}
}

func main() {
	fmt.Println(myPow(2, 6))
	fmt.Println(myPow(2, 0))
	fmt.Println(myPow(2, -1))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(2, -3))
	fmt.Println(myPow(0.00001, 2147483647))
}
