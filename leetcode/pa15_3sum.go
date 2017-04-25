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
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	for i:=0; i<len(nums)-2;i++ {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		target:= nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left] + nums[right] + target == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left] + nums[right] + target > 0 {
				right--
			} else {
				left++
			}
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
