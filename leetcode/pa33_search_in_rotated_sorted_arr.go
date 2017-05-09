package main

import "fmt"

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.


*/
func search(nums []int, target int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return -1
	}
	half := len(nums) / 2
	if nums[half] == target {
		return half
	}
	if nums[0] < nums[len(nums)-1] {
		return binSearch(nums, target)
	}
	if nums[half] < nums[0] {
		// half is in small rotated.
		if nums[len(nums)-1] < target {
			return search(nums[:half], target)
		} else {
			if nums[half] < target {
				pos := binSearch(nums[half:], target)
				if pos != -1 {
					pos += half
				}
				return pos
			} else {
				return search(nums[:half], target)
			}
		}
	} else {
		// bigger rotated
		if nums[0] > target {
			pos := search(nums[half:], target)
			if pos != -1 {
				pos += half
			}
			return pos
		} else {
			if nums[half] > target {
				return search(nums[0:half], target)
			} else {
				pos := search(nums[half:], target)
				if pos != -1 {
					pos += half
				}
				return pos
			}
		}
	}
}

func binSearch(nums []int, target int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return -1
	}
	half := len(nums) / 2
	if nums[half] == target {
		return half
	}
	if nums[half] > target {
		return binSearch(nums[:half], target)
	} else {
		pos:= binSearch(nums[half:], target)
		if pos != -1 {
			pos += half
		}
		return pos
	}
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 3}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 3}, 7))
	fmt.Println(search([]int{1, 3, 5, 7, 8}, 7))
	fmt.Println(search([]int{7, 8, 9, 1, 2, 4, 5, 6}, 3))
	fmt.Println(search([]int{7, 8, 9, 1, 2, 4, 5, 6}, 3))
	fmt.Println(search([]int{2,3, 4, 5, 6, 7, 8, 9, 1}, 9))
}
