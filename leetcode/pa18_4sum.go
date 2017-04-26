package main

import "fmt"

/*
For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */
func fourSum(nums []int, target int) [][]int {
	sort(nums, 0, len(nums)-1)
	if len(nums) < 4 {
		return [][]int{}
	}
	result:=[][]int{}
	for i:=0; i<len(nums)-3;i++  {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		tr:=threeSum(nums[i+1:], target-nums[i])
		for _, r:= range tr {
			result=append(result, []int{nums[i], r[0], r[1], r[2]})
		}
	}
	return result
}

func threeSum(nums []int, t int) [][]int {
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
			if nums[left] + nums[right] + target == t {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left] + nums[right] + target > t {
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
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
