package main

import (
	"fmt"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -1 * n }
func product(m, n int) int { return m * n }

// 词法作用域
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func sum(vals...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3))
	fp := product
	fmt.Println(fp(2, f(3)))

	// 变量 x 在 p中。
	p := squares()

	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())

	fmt.Println(squares()())

	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	values := []int{10, 20, 30, 40}
	fmt.Println(sum(values...))
}
