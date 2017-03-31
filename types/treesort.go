package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t==nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int  {
	if t!=nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

func main()  {
	arr := []int{1,2,34,23,1111,3334,434}
	fmt.Println("before:", arr)
	fmt.Println("after:", Sort(arr))
}
