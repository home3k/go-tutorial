package main

import "fmt"

// 整个package可见
const ca = 1000

func main()  {

	var local = 100

	// init -> s = ""
	var s string

	fmt.Printf("%d %s \n", local, s)

	var i, j, k int
	var a, b, c = true, "x", 3
	fmt.Println(i, j, k, a, b, c)

	p := &c
	fmt.Println(*p)
	fmt.Println(p)
	*p = 300
	fmt.Println(c)

	var x, y int
	fmt.Println(x==y, &x==&x, &x==&y, &x==nil)

	ni := new(int)
	fmt.Println(ni, *ni)
	*ni = 200
	fmt.Println(ni, *ni)

	var w1, w2 = 1, 2
	fmt.Println(w1, w2)
	w1, w2 = w2, w1
	fmt.Println(w1, w2)
}
