package main

import "fmt"

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

For example,
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
                |
       |        |
   |   ||  |  | |
_|_||_|||| ||||||
0123456789 0123456
*/

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	result := 0
	stack := []int{}
	minus := 0
	minHeight := 0
	for i, h := range height {
		if h <= minHeight {
			if len(stack) > 0 {
				minus += h
			}
			continue
		}
		if len(stack) == 0 {
			minHeight = h
			stack = append(stack, i)
			minus = 0
		} else {
			top := stack[0]
			if i > top+1 {
				result += min(h, height[top]) * (i - top - 1)
				stack = []int{}
				minHeight = h
				result -= minus
				stack = append(stack, i)
				minus = 0
			} else {
				stack = []int{}
				minHeight = h
				stack = append(stack, i)
			}

		}

	}
	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 1, 2, 1, 4}))
}
