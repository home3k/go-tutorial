package main

import (
	"fmt"
)

/*
Given a set of candidate numbers (C) (without duplicates) and a target number (T),
find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7,

A solution set is:
[
  [7],
  [2, 2, 3]
]
*/
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	result := [][]int{}
	sort(candidates, 0, len(candidates)-1)
	if search(candidates, target) {
		result = append(result, []int{target})
	}
	start, end := 0, len(candidates)
	for start < end {
		val := candidates[start]
		if val > target {
			break
		}
		i := 0
		tVal := val
		tresult:=[][]int{}
		for tVal < target {
			i++
			if target-tVal < val {
				break
			}
			tv := []int{}
			for x := 0; x < i; x++ {
				tv = append(tv, val)
			}
			if start+1 < end && target-tVal >= candidates[start+1] {
				next := combinationSum(candidates[start+1:], target-tVal)
				for x:=len(next)-1;x>=0;x-- {
					ns:=next[x]
					tresult = append([][]int{append(tv, ns...)}, tresult...)
				}
			}
			if target-tVal == val {
				tv = append(tv, val)
				tresult = append([][]int{tv}, tresult...)
				break
			}
			tVal += val
		}
		result=append(result, tresult...)
		start++
	}
	return result
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}
	half := len(nums) / 2
	if nums[half] == target {
		return true
	} else if nums[half] > target {
		return search(nums[:half], target)
	} else {
		return search(nums[half+1:], target)
	}
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
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{7, 3, 6, 2}, 7))
	fmt.Println(combinationSum([]int{2, 3, 6, 1}, 7))
}
