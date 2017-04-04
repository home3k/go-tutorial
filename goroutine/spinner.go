package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main()  {
	spinner(100*time.Millisecond)
	const N  =  45
	fibN := fib(N)
	fmt.Printf("fibonacci(%d) result: %d\n", N, fibN)
}
