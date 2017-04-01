package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	var m = make(map[int]int)
	for index, num := range nums {
		// match
		if val, ok := m[num]; ok {
			result = []int{val, index}
			break
		}
		m[target-num] = index
	}
	return result
}

func main() {
	// test
	fmt.Println(twoSum([]int{1, 5, 8}, 6))
	fmt.Println(twoSum([]int{10, 5, 8, 28, 1}, 9))
}
