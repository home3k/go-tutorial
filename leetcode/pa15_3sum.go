package main

import "fmt"

/*

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
-4 -1 -1 0 1 2
-2
-3 -1 -1 0 2 2
-2 1 2

*/
func threeSum(nums []int) [][]int {
	sort(nums, 0, len(nums)-1)
	left, right := 0, len(nums)-1
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	for left < right-1 {
		index := left + 1
		for index < right {
			sum := nums[left] + nums[right] + nums[index]
			if sum == 0 {
				result = append(result, []int{nums[left], nums[index], nums[right]})
				index++
			} else if sum > 0 {
				break
			} else {
				index++
			}
		}
		if nums[left]+nums[right] <0 {
			left++
		} else {
			right--
		}
	}
	return result
}

func sort(array []int, start int, end int) {
	if start >= end {
		return
	}
	pivot := array[(start+end)/2]
	left, right := start, end
	for left <= right {
		for left <= right && array[left] < pivot {
			left++
		}
		for left <= right && array[right] > pivot {
			right--
		}
		if left <= right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
	}
	sort(array, start, right)
	sort(array, left, end)
}

func main()  {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
