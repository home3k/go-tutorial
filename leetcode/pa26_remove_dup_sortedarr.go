package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	idx := 1
	for i:= 1; i<len(nums);i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func main() {
	var c = []int{1, 2, 3, 4}
	fmt.Println(removeDuplicates(c), c)
	c = []int{1, 3, 3}
	fmt.Println(removeDuplicates(c), c)
	c = []int{1, 1, 3}
	fmt.Println(removeDuplicates(c), c)
	c = []int{}
	fmt.Println(removeDuplicates(c), c)
	c = []int{1, 1, 2, 3, 3, 4, 5, 6, 6}
	fmt.Println(removeDuplicates(c), c)
}