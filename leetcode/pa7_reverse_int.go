package main

import "fmt"

const MAX = 1<<31 - 1

func reverse(x int) int {
	positive := true
	if x < 0 {
		positive = false
		x = x * -1
	}
	if x < 10 {
		return x
	}
	ne := 0
	for x > 0 {
		ne = ne*10 + x%10
		x = x / 10
	}
	if ne > MAX || ne < (-1*MAX) {
		return 0
	}
	if !positive {
		return -1 * ne
	}
	return ne
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(0))
	fmt.Println(reverse(10))
	fmt.Println(reverse(20100))
	fmt.Println(reverse(1000000003))
}
