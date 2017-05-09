package main

import "fmt"

/*
Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].


*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return []int{-1, -1}
	}
	result := []int{}
	half := (len(nums) - 1) / 2
	if nums[half] == target {
		left := _search(nums, target, 0, half-1, true)
		right := _search(nums, target, half+1, len(nums)-1, false)
		if left != -1 && left < half {
			result = append(result, left)
		} else {
			result = append(result, half)
		}
		if right != -1 && right > half {
			result = append(result, right)
		} else {
			result = append(result, half)
		}
	} else if nums[half] > target {
		return searchRange(nums[:half], target)
	} else {
		p := searchRange(nums[half+1:], target)
		if p[0] == -1 {
			return p
		}
		for i := range p {
			p[i] += half + 1
		}
		return p
	}
	return result
}

func _search(nums []int, target int, start int, end int, left bool) int {
	if start > end {
		return -1
	}
	if start == end {
		if target == nums[start] {
			return start
		} else {
			return -1
		}
	}
	half := (start + end) / 2
	if nums[half] == target {
		var ns, ne int
		if left {
			ns, ne = start, half-1
		} else {
			ns, ne = half+1, end
		}
		v := _search(nums, target, ns, ne, left)
		if v == -1 {
			return half
		} else {
			return v
		}
	} else {
		// 5 , 7, 7, 8, 8, 10
		if left && nums[half] < target {
			return _search(nums, target, half+1, end, left)
		} else if !left && nums[half] > target {
			return _search(nums, target, start, half-1, left)
		} else {
			// can not happen
			return -1
		}
	}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 9))
	fmt.Println(searchRange([]int{5}, 9))
	fmt.Println(searchRange([]int{}, 9))
	fmt.Println(searchRange([]int{8,8,8,8,8,8,8}, 8))
	fmt.Println(searchRange([]int{7,8,8,8,8,8,9}, 8))
}
