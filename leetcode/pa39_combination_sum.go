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
	return _combinationSum(candidates, target, true)
}

func _combinationSum(candidates []int, target int, _sort bool )  [][]int  {
	result:=[][]int{}
	if len(candidates) == 0 {
		return result
	}
	if _sort {
		sort(candidates, 0, len(candidates)-1)
	}
	start, end := 0, len(candidates)
	for start < end {
		val := candidates[start]
		if val > target {
			break
		}
		tVal := target
		if target % val == 0 {
			tv := []int{}
			for x := 0; x < target / val; x++ {
				tv = append(tv, val)
			}
			result = append(result, tv)
			tVal = target - 2 * val
		} else {
			tVal = target - val
		}
		for tVal >= val {

			i:=tVal / val
			rest:=target - i*val
			if rest < val {
				break
			}

			if start+1 < end && rest >= candidates[start+1] {
				next := _combinationSum(candidates[start+1:], rest, false)
				for _, ns:=range next {
					tv := []int{}
					for x := 0; x < i; x++ {
						tv = append(tv, val)
					}
					tv=append(tv,ns...)
					result = append(result, tv)
					fmt.Println(result)
				}
			}
			tVal -= val
		}
		start++
	}
	//if search(candidates, target) {
	//	result = append(result, []int{target})
	//}
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
	fmt.Println(combinationSum([]int{5, 10, 8, 4, 3, 12, 9}, 27))
}
