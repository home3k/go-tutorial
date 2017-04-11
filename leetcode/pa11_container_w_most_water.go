package main

import "fmt"

func maxArea(height []int) int {
	left, right, result := 0, len(height)-1, 0
	for left < right {
		result = max(result, -max(-height[left], -height[right])*(right-left))
		if height[left]<height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func max(v,f int) int {
	if v>f {
		return v
	}
	return f
}

func main() {
	fmt.Println(maxArea([]int{1, 2, 3, 5}))
	fmt.Println(maxArea([]int{8, 10, 12, 5}))
}
